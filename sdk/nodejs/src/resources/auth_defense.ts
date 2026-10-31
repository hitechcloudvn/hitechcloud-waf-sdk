import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Dynamic Authentication Defense: Authentication-defense sources, protected users, reviews and records. */
export class AuthDefenseResource extends BaseResource {
  /** GET /open/auth_defense/source - List authentication-defense sources */
  listSources(params?: Params): Promise<any> {
    return this.http.get("/open/auth_defense/source", params);
  }
  /** POST /open/auth_defense/source - Create a source */
  createSource(data?: Body): Promise<any> {
    return this.http.post("/open/auth_defense/source", data);
  }
  /** GET /open/auth_defense/source/password - Get the shared source password policy */
  getSourcePassword(params?: Params): Promise<any> {
    return this.http.get("/open/auth_defense/source/password", params);
  }
  /** PUT /open/auth_defense/source/password - Update the shared source password policy */
  updateSourcePassword(data?: Body): Promise<any> {
    return this.http.put("/open/auth_defense/source/password", data);
  }
  /** GET /open/auth_defense/source/{id} - Get a source */
  getSource(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/auth_defense/source/${id}`, params);
  }
  /** PUT /open/auth_defense/source/{id} - Update a source */
  updateSource(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/auth_defense/source/${id}`, data);
  }
  /** DELETE /open/auth_defense/source/{id} - Delete a source */
  deleteSource(id: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/auth_defense/source/${id}`, data);
  }
  /** GET /open/auth_defense/source/{id}/secret - Get a source secret */
  getSourceSecret(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/auth_defense/source/${id}/secret`, params);
  }
  /** GET /open/auth_defense/source/{id}/user - List users bound to a source */
  listSourceUsers(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/auth_defense/source/${id}/user`, params);
  }
  /** PUT /open/auth_defense/source/{id}/user/{user_id} - Update a source user */
  updateSourceUser(id: string, userId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/auth_defense/source/${id}/user/${userId}`, data);
  }
  /** DELETE /open/auth_defense/source/{id}/user/{user_id} - Remove a user from a source */
  deleteSourceUser(id: string, userId: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/auth_defense/source/${id}/user/${userId}`, data);
  }
  /** GET /open/auth_defense/user - List protected users */
  listUsers(params?: Params): Promise<any> {
    return this.http.get("/open/auth_defense/user", params);
  }
  /** POST /open/auth_defense/user - Create a protected user */
  createUser(data?: Body): Promise<any> {
    return this.http.post("/open/auth_defense/user", data);
  }
  /** DELETE /open/auth_defense/user - Delete protected users (bulk) */
  deleteUsers(data?: Body): Promise<any> {
    return this.http.delete("/open/auth_defense/user", data);
  }
  /** POST /open/auth_defense/user/merge - Merge protected users */
  mergeUsers(data?: Body): Promise<any> {
    return this.http.post("/open/auth_defense/user/merge", data);
  }
  /** GET /open/auth_defense/user/review - List pending user reviews */
  listUserReviews(params?: Params): Promise<any> {
    return this.http.get("/open/auth_defense/user/review", params);
  }
  /** PUT /open/auth_defense/user/review/{review_id} - Approve / reject a user review */
  updateUserReview(reviewId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/auth_defense/user/review/${reviewId}`, data);
  }
  /** GET /open/auth_defense/user/{user_id} - Get a protected user */
  getUser(userId: string, params?: Params): Promise<any> {
    return this.http.get(`/open/auth_defense/user/${userId}`, params);
  }
  /** PUT /open/auth_defense/user/{user_id} - Update a protected user */
  updateUser(userId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/auth_defense/user/${userId}`, data);
  }
  /** DELETE /open/auth_defense/user/{user_id} - Delete a protected user */
  deleteUser(userId: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/auth_defense/user/${userId}`, data);
  }
  /** PUT /open/auth_defense/user/{user_id}/reset_totp - Reset a user's TOTP */
  resetUserTotp(userId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/auth_defense/user/${userId}/reset_totp`, data);
  }
  /** PUT /open/auth_defense/user/{user_id}/unbind - Unbind a user's authenticator */
  unbindUser(userId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/auth_defense/user/${userId}/unbind`, data);
  }
  /** GET /open/records/auth_defense - Authentication-defense event records */
  records(params?: Params): Promise<any> {
    return this.http.get("/open/records/auth_defense", params);
  }
  /** GET /open/v2/records/auth_defense - Authentication-defense event records (v2) */
  recordsV2(params?: Params): Promise<any> {
    return this.http.get("/open/v2/records/auth_defense", params);
  }
  /** POST /business/users/{id}/unlock - Unlock a locked business user */
  unlockUser(id: string, data?: Body): Promise<any> {
    return this.http.post(`/business/users/${id}/unlock`, data);
  }
}
