import axios, { AxiosInstance, AxiosError } from "axios";
import {
  WafError,
  AuthenticationError,
  AuthorizationError,
  NotFoundError,
  ValidationError,
  RateLimitError,
  ServerError,
} from "./errors";

export type Params = Record<string, unknown>;
export type Body = Record<string, unknown>;

export class HttpClient {
  private client: AxiosInstance;
  private maxRetries: number;
  token: string | null = null;

  constructor(baseUrl: string, timeout = 30000, maxRetries = 3) {
    this.maxRetries = maxRetries;
    this.client = axios.create({
      baseURL: baseUrl.replace(/\/+$/, "") + "/api",
      timeout,
      headers: {
        "Content-Type": "application/json",
        "User-Agent": "HiTechCloud-WAF-NodeJS/1.0.0",
      },
    });
  }

  setToken(token: string): void {
    this.token = token;
    this.client.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  }

  clearToken(): void {
    this.token = null;
    delete this.client.defaults.headers.common["Authorization"];
  }

  private mapError(error: AxiosError): never {
    if (!error.response) {
      throw new WafError(error.message);
    }
    const { status, data } = error.response as { status: number; data: any };
    const message = (data && (data.msg || data.err)) || `HTTP ${status}`;
    switch (status) {
      case 401:
        throw new AuthenticationError(message, data);
      case 403:
        throw new AuthorizationError(message, data);
      case 404:
        throw new NotFoundError(message, data);
      case 422:
        throw new ValidationError(message, data);
      case 429: {
        const retryAfter = parseInt(
          String((error.response.headers as any)["retry-after"] || "60"), 10);
        throw new RateLimitError(message, retryAfter, data);
      }
      default:
        if (status >= 500) throw new ServerError(message, status, data);
        throw new WafError(message, status, data);
    }
  }

  async request<T = any>(method: string, path: string, data?: Body | null, params?: Params | null): Promise<T> {
    let lastError: AxiosError | undefined;
    for (let attempt = 0; attempt <= this.maxRetries; attempt++) {
      try {
        const resp = await this.client.request<T>({ method, url: path, data: data ?? undefined, params: params ?? undefined });
        return resp.data;
      } catch (err) {
        const error = err as AxiosError;
        lastError = error;
        const status = error.response?.status;
        if (status === 429 && attempt < this.maxRetries) {
          const retryAfter = parseInt(String((error.response!.headers as any)["retry-after"] || "60"), 10);
          await new Promise((r) => setTimeout(r, retryAfter * 1000));
          continue;
        }
        if (attempt < this.maxRetries && (!error.response || (status && status >= 500))) {
          await new Promise((r) => setTimeout(r, Math.pow(2, attempt) * 1000));
          continue;
        }
        this.mapError(error);
      }
    }
    this.mapError(lastError as AxiosError);
  }

  get<T = any>(path: string, params?: Params | null): Promise<T> {
    return this.request<T>("GET", path, null, params);
  }
  post<T = any>(path: string, data?: Body | null): Promise<T> {
    return this.request<T>("POST", path, data);
  }
  put<T = any>(path: string, data?: Body | null): Promise<T> {
    return this.request<T>("PUT", path, data);
  }
  delete<T = any>(path: string, data?: Body | null): Promise<T> {
    return this.request<T>("DELETE", path, data);
  }
}
