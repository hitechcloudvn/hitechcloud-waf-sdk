import { HttpClient } from "../http-client";

export class BaseResource {
  protected http: HttpClient;
  constructor(http: HttpClient) {
    this.http = http;
  }
}
