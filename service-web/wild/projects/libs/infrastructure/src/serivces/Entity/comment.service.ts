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
export class CommentService {
  siteUrl = `v1/comment`;

  constructor(private http: HttpClient) {}

  Add(model: CommentModels.AddRequest): Observable<CommentModels.AddResponse> {
    return this.http.post<CommentModels.AddResponse>(
      `${this.siteUrl}/add`,
      model
    );
  }

  Like(
    model: CommentModels.LikeRequest
  ): Observable<CommentModels.LikeResponse> {
    return this.http.post<CommentModels.LikeResponse>(
      `${this.siteUrl}/like`,
      model
    );
  }

  List(
    model: CommentModels.ListRequest
  ): Observable<CommentModels.ListResponse> {
    return this.http.post<CommentModels.ListResponse>(
      `${this.siteUrl}/list`,
      model
    );
  }

  MyComments(
    model: CommentModels.MyCommentsRequest
  ): Observable<CommentModels.MyCommentsResponse> {
    return this.http.post<CommentModels.MyCommentsResponse>(
      `${this.siteUrl}/my-comments`,
      model
    );
  }
}

export namespace CommentModels {
  export class AddRequest {
    public Comment!: AppModels.CommentModel;
  }

  export class AddResponse {
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
    public PostID?: string;
  }

  export class ListResponse {
    public Comments!: AppModels.CommentModel[];
    public ExistMore!: boolean;
  }

  export class MyCommentsRequest {
    public Page!: number;
  }

  export class MyCommentsResponse {
    public Comments!: AppModels.CommentModel[];
    public ExistMore!: boolean;
  }
}
