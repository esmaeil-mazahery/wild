import { AppStorage, StorageKeys } from '../../utils/localStorage';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs/operators';
import { BehaviorSubject, Observable } from 'rxjs';
import { Extention } from '../../utils/generic';
import { BuildHttpParamsWithModel } from '../utils';
import { AppModels } from './models';

@Injectable({
  providedIn: 'root',
})
export class AuthenticationService {
  siteUrl = `v1/auth`;
  invalidLogin = true;

  private updateProfile = new BehaviorSubject<AppModels.MemberModel>(
    this.GetProfile()
  );

  constructor(private http: HttpClient) {}

  GetProfileSync(): Observable<AppModels.MemberModel> {
    return this.updateProfile.asObservable();
  }

  UpdateProfile(profile?: AppModels.MemberModel) {
    if (profile != null) {
      AppStorage.setItem(StorageKeys.Token, profile.Token);
      AppStorage.setItem(StorageKeys.Follower, profile.Follower);
      AppStorage.setItem(StorageKeys.Following, profile.Following);
      AppStorage.setItem(StorageKeys.Biography, profile.Biography ?? '');
      AppStorage.setItem(StorageKeys.Username, profile.Username);
      AppStorage.setItem(
        StorageKeys.Mobile,
        Extention.ConvertString(profile.Mobile)
      );
      AppStorage.setItem(
        StorageKeys.Email,
        Extention.ConvertString(profile.Email)
      );
      AppStorage.setItem(StorageKeys.Image, profile.Image ?? '');
      AppStorage.setItem(StorageKeys.ImageHeader, profile.ImageHeader ?? '');
      AppStorage.setItem(
        StorageKeys.Name,
        Extention.ConvertString(profile.Name)
      );
      AppStorage.setItem(
        StorageKeys.Family,
        Extention.ConvertString(profile.Family)
      );
    }
    this.updateProfile.next(profile || this.GetProfile());
  }

  IsLogin() {
    const token = AppStorage.getItem(StorageKeys.Token);
    if (token) {
      return true;
    }
    return false;
  }

  GetProfile(): AppModels.MemberModel {
    return {
      Username: AppStorage.getItem(StorageKeys.Username),
      Email: AppStorage.getItem(StorageKeys.Email),
      Family: AppStorage.getItem(StorageKeys.Family),
      Image: AppStorage.getItem(StorageKeys.Image),
      Mobile: AppStorage.getItem(StorageKeys.Mobile),
      Name: AppStorage.getItem(StorageKeys.Name),
      Token: AppStorage.getItem(StorageKeys.Token),
      ImageHeader: AppStorage.getItem(StorageKeys.ImageHeader),
      Follower: AppStorage.getItem(StorageKeys.Follower),
      Following: AppStorage.getItem(StorageKeys.Following),
      Biography: AppStorage.getItem(StorageKeys.Biography),
    };
  }

  Login(
    model: AuthenticationModels.LoginRequest
  ): Observable<AuthenticationModels.LoginResponse> {
    return this.http
      .post<AuthenticationModels.LoginResponse>(`${this.siteUrl}/login`, model)
      .pipe(
        tap((response) => {
          this.UpdateProfile({
            Image: response.Member.Image,
            Username: response.Member.Username,
            Name: response.Member.Name,
            Family: response.Member.Family,
            Email: response.Member.Email,
            Mobile: response.Member.Mobile,
            Token: response.Member.Token,
            ImageHeader: response.Member.ImageHeader,
            Following: response.Member.Following,
            Follower: response.Member.Follower,
            Biography: response.Member.Biography,
          });
        })
      );
  }

  LogOut() {
    AppStorage.clear();
  }

  ForgetPassword(
    model: AuthenticationModels.ForgetPasswordRequest
  ): Observable<AuthenticationModels.ForgetPasswordResponse> {
    return this.http.post<AuthenticationModels.ForgetPasswordResponse>(
      `${this.siteUrl}/forget-password`,
      model
    );
  }

  ForgetPasswordChange(
    model: AuthenticationModels.ForgetPasswordChangeRequest
  ): Observable<AuthenticationModels.ForgetPasswordChangeResponse> {
    return this.http.post<AuthenticationModels.ForgetPasswordChangeResponse>(
      `${this.siteUrl}/forget-password-change`,
      model
    );
  }

  Register(
    model: AuthenticationModels.RegisterRequest
  ): Observable<AuthenticationModels.RegisterResponse> {
    return this.http
      .post<AuthenticationModels.RegisterResponse>(
        `${this.siteUrl}/register`,
        model
      )
      .pipe(
        tap((response) => {
          this.UpdateProfile({
            Image: response.Member.Image,
            Username: response.Member.Username,
            Name: response.Member.Name,
            Family: response.Member.Family,
            Email: response.Member.Email,
            Mobile: response.Member.Mobile,
            Token: response.Member.Token,
          });
        })
      );
  }

  ChangeImageProfile(
    model: AuthenticationModels.ChangeImageProfileRequest
  ): Observable<AuthenticationModels.ChangeImageProfileResponse> {
    return this.http
      .post<AuthenticationModels.ChangeImageProfileResponse>(
        `${this.siteUrl}/change-image-profile`,
        model
      )
      .pipe(
        tap(() => {
          this.UpdateProfile({
            ...this.GetProfile(),
            Image: model.ImageURL,
          });
        })
      );
  }

  ChangeImageHeader(
    model: AuthenticationModels.ChangeImageHeaderRequest
  ): Observable<AuthenticationModels.ChangeImageHeaderResponse> {
    return this.http
      .post<AuthenticationModels.ChangeImageHeaderResponse>(
        `${this.siteUrl}/change-image-header`,
        model
      )
      .pipe(
        tap(() => {
          this.UpdateProfile({
            ...this.GetProfile(),
            ImageHeader: model.ImageURL,
          });
        })
      );
  }

  ChangePassword(
    model: AuthenticationModels.ChangePasswordRequest
  ): Observable<AuthenticationModels.ChangePasswordResponse> {
    return this.http.post<AuthenticationModels.ChangePasswordResponse>(
      `${this.siteUrl}/change-password`,
      model
    );
  }

  Suggestion(
    model: AuthenticationModels.SuggestionRequest
  ): Observable<AuthenticationModels.SuggestionResponse> {
    return this.http.post<AuthenticationModels.SuggestionResponse>(
      `${this.siteUrl}/suggestion`,
      model
    );
  }

  Followers(
    model: AuthenticationModels.FollowersRequest
  ): Observable<AuthenticationModels.FollowersResponse> {
    return this.http.post<AuthenticationModels.FollowersResponse>(
      `${this.siteUrl}/followers`,
      model
    );
  }

  Followings(
    model: AuthenticationModels.FollowingsRequest
  ): Observable<AuthenticationModels.FollowingsResponse> {
    return this.http.post<AuthenticationModels.FollowingsResponse>(
      `${this.siteUrl}/followings`,
      model
    );
  }

  Follow(
    model: AuthenticationModels.FollowRequest
  ): Observable<AuthenticationModels.FollowResponse> {
    return this.http.post<AuthenticationModels.FollowResponse>(
      `${this.siteUrl}/follow`,
      model
    );
  }
}

export namespace AuthenticationModels {
  export class LoginRequest {
    public Username: string | undefined;
    public Password: string | undefined;
  }

  export class LoginResponse {
    public Member!: AppModels.MemberModel;
  }

  export class ForgetPasswordRequest {
    public Username: string | undefined;
  }

  export class ForgetPasswordResponse {
    public Token: string | undefined;
  }

  export class ForgetPasswordChangeRequest {
    public Token: string | undefined;
    public Password: string | undefined;
  }

  export class ForgetPasswordChangeResponse {
    public Member: AppModels.MemberModel | undefined;
  }

  export class RegisterRequest {
    public Member: AppModels.MemberModel | undefined;
  }

  export class RegisterResponse {
    public Member!: AppModels.MemberModel;
  }

  export class ChangeImageProfileRequest {
    public ImageURL: string | undefined;
  }

  export class ChangeImageProfileResponse {}

  export class ChangeImageHeaderRequest {
    public ImageURL: string | undefined;
  }

  export class ChangeImageHeaderResponse {}

  export class ChangePasswordRequest {
    public Username: string | undefined;
    public Password: string | undefined;
    public NewPassword: string | undefined;
  }

  export class ChangePasswordResponse {
    public Member: AppModels.MemberModel | undefined;
  }

  export class SuggestionRequest {}

  export class SuggestionResponse {
    public Members: AppModels.MemberModel[] = [];
  }

  export class FollowersRequest {}

  export class FollowersResponse {
    public Members: AppModels.MemberModel[] = [];
  }

  export class FollowingsRequest {}

  export class FollowingsResponse {
    public Members: AppModels.MemberModel[] = [];
  }

  export class FollowRequest {
    public MemberID!: string;
    public Follow!: boolean;
  }

  export class FollowResponse {
    public Result!: boolean;
  }
}
