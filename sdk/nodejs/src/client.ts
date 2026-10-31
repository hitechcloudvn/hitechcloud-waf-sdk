import { HttpClient } from "./http-client";
import { AuthResource } from "./resources/auth";
import { UsersResource } from "./resources/users";
import { AuthDefenseResource } from "./resources/auth_defense";
import { SitesResource } from "./resources/sites";
import { PoliciesResource } from "./resources/policies";
import { SkynetResource } from "./resources/skynet";
import { AntiTamperResource } from "./resources/anti_tamper";
import { AclResource } from "./resources/acl";
import { ChallengeResource } from "./resources/challenge";
import { WaitingRoomResource } from "./resources/waiting_room";
import { Ja4Resource } from "./resources/ja4";
import { GlobalSettingsResource } from "./resources/global_settings";
import { CloudResource } from "./resources/cloud";
import { SslResource } from "./resources/ssl";
import { EventsResource } from "./resources/events";
import { SecurityPostureResource } from "./resources/security_posture";
import { StatisticsResource } from "./resources/statistics";
import { ReportResource } from "./resources/report";
import { AlarmResource } from "./resources/alarm";
import { PortalResource } from "./resources/portal";
import { AntiBotResource } from "./resources/anti_bot";
import { SettingsResource } from "./resources/settings";
import { SystemResource } from "./resources/system";
import { LicenseResource } from "./resources/license";
import { McpResource } from "./resources/mcp";

export interface WafClientOptions {
  baseUrl?: string;
  username?: string;
  password?: string;
  token?: string;
  tfaCode?: string;
  sandbox?: boolean;
  timeout?: number;
  maxRetries?: number;
}

const PRODUCTION_URL = "https://vcwaf.hitechcloud.vn";
const SANDBOX_URL = "https://vcwaf-sandbox.hitechcloud.vn";

function dig(obj: any, key: string): any {
  if (!obj || typeof obj !== "object") return undefined;
  if (key in obj) return obj[key];
  if (obj.data && typeof obj.data === "object") return obj.data[key];
  return undefined;
}

export class WafClient {
  private http: HttpClient;
  readonly auth: AuthResource;
  readonly users: UsersResource;
  readonly authDefense: AuthDefenseResource;
  readonly sites: SitesResource;
  readonly policies: PoliciesResource;
  readonly skynet: SkynetResource;
  readonly antiTamper: AntiTamperResource;
  readonly acl: AclResource;
  readonly challenge: ChallengeResource;
  readonly waitingRoom: WaitingRoomResource;
  readonly ja4: Ja4Resource;
  readonly globalSettings: GlobalSettingsResource;
  readonly cloud: CloudResource;
  readonly ssl: SslResource;
  readonly events: EventsResource;
  readonly securityPosture: SecurityPostureResource;
  readonly statistics: StatisticsResource;
  readonly report: ReportResource;
  readonly alarm: AlarmResource;
  readonly portal: PortalResource;
  readonly antiBot: AntiBotResource;
  readonly settings: SettingsResource;
  readonly system: SystemResource;
  readonly license: LicenseResource;
  readonly mcp: McpResource;

  constructor(options: WafClientOptions = {}) {
    const baseUrl = options.baseUrl || (options.sandbox ? SANDBOX_URL : PRODUCTION_URL);
    this.http = new HttpClient(baseUrl, (options.timeout ?? 30) * 1000, options.maxRetries ?? 3);

    this.auth = new AuthResource(this.http);
    this.users = new UsersResource(this.http);
    this.authDefense = new AuthDefenseResource(this.http);
    this.sites = new SitesResource(this.http);
    this.policies = new PoliciesResource(this.http);
    this.skynet = new SkynetResource(this.http);
    this.antiTamper = new AntiTamperResource(this.http);
    this.acl = new AclResource(this.http);
    this.challenge = new ChallengeResource(this.http);
    this.waitingRoom = new WaitingRoomResource(this.http);
    this.ja4 = new Ja4Resource(this.http);
    this.globalSettings = new GlobalSettingsResource(this.http);
    this.cloud = new CloudResource(this.http);
    this.ssl = new SslResource(this.http);
    this.events = new EventsResource(this.http);
    this.securityPosture = new SecurityPostureResource(this.http);
    this.statistics = new StatisticsResource(this.http);
    this.report = new ReportResource(this.http);
    this.alarm = new AlarmResource(this.http);
    this.portal = new PortalResource(this.http);
    this.antiBot = new AntiBotResource(this.http);
    this.settings = new SettingsResource(this.http);
    this.system = new SystemResource(this.http);
    this.license = new LicenseResource(this.http);
    this.mcp = new McpResource(this.http);

    if (options.token) {
      this.http.setToken(options.token);
    }
  }

  /** Run the full CSRF -> login (-> TFA) flow and store the JWT. */
  async login(username: string, password: string, tfaCode?: string): Promise<any> {
    const csrf = (await this.auth.csrf()) || {};
    const csrfToken = dig(csrf, "csrf_token") || dig(csrf, "token");
    const payload: Record<string, unknown> = { username, password };
    if (csrfToken) payload.csrf_token = csrfToken;
    let result = (await this.auth.login(payload)) || {};
    let token = dig(result, "jwt") || dig(result, "token");
    if (tfaCode && !token) {
      result = (await this.auth.tfa({ code: tfaCode, username })) || {};
      token = dig(result, "jwt") || dig(result, "token");
    }
    if (token) this.http.setToken(token);
    return result;
  }

  setToken(token: string): void {
    this.http.setToken(token);
  }

  async logout(): Promise<any> {
    const result = await this.auth.logout();
    this.http.clearToken();
    return result;
  }
}
