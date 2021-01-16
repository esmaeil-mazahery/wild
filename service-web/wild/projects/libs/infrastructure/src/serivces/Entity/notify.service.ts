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
export class NotifyService {
  siteUrl = `v1/notify`;

  constructor(private http: HttpClient) {}

  List(model: PostModels.ListRequest): Observable<PostModels.ListResponse> {
    return this.http.post<PostModels.ListResponse>(
      `${this.siteUrl}/list`,
      model
    );
  }

  Read(model: PostModels.ReadRequest): Observable<PostModels.ReadResponse> {
    return this.http.post<PostModels.ReadResponse>(
      `${this.siteUrl}/read`,
      model
    );
  }
}

export namespace PostModels {
  export class ListRequest {
    public Page!: number;
  }

  export class ListResponse {
    public Notifies!: AppModels.NotifyModel[];
    public ExistMore!: boolean;
  }

  export class ReadRequest {
    public IDs: string[] = [];
  }

  export class ReadResponse {}
}
