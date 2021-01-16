import { startWith } from 'rxjs/operators';
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
export class PostService {
  siteUrl = `v1/post`;

  constructor(private http: HttpClient) {}

  Register(
    model: PostModels.RegisterRequest
  ): Observable<PostModels.RegisterResponse> {
    return this.http.post<PostModels.RegisterResponse>(
      `${this.siteUrl}/register`,
      model
    );
  }

  Like(model: PostModels.LikeRequest): Observable<PostModels.LikeResponse> {
    return this.http.post<PostModels.LikeResponse>(
      `${this.siteUrl}/like`,
      model
    );
  }

  List(model: PostModels.ListRequest): Observable<PostModels.ListResponse> {
    return this.http.post<PostModels.ListResponse>(
      `${this.siteUrl}/list`,
      model
    );
  }

  MyPosts(
    model: PostModels.MyPostsRequest
  ): Observable<PostModels.MyPostsResponse> {
    return this.http.post<PostModels.MyPostsResponse>(
      `${this.siteUrl}/my-posts`,
      model
    );
  }

  Search(
    model: PostModels.SearchRequest
  ): Observable<PostModels.SearchResponse> {
    return this.http.post<PostModels.SearchResponse>(
      `${this.siteUrl}/search`,
      model
    );
  }
}

export namespace PostModels {
  export class RegisterRequest {
    public Post!: AppModels.PostModel;
  }

  export class RegisterResponse {
    public ID!: string;
  }

  export class LikeRequest {
    public ID!: string;
    public Like!: boolean;
  }

  export class LikeResponse {
    public Result!: boolean;
  }

  export class ListRequest {
    public Page!: number;
  }

  export class ListResponse {
    public Posts!: AppModels.PostModel[];
    public ExistMore!: boolean;
  }

  export class MyPostsRequest {
    public Page!: number;
  }

  export class MyPostsResponse {
    public Posts!: AppModels.PostModel[];
    public ExistMore!: boolean;
  }

  export class SearchRequest {
    public Page!: number;
    public Term!: string;
  }

  export class SearchResponse {
    public Posts!: AppModels.PostModel[];
    public ExistMore!: boolean;
  }
}
