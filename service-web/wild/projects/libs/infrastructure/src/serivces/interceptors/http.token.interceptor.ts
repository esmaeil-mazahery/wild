import { AppStorage, StorageKeys } from './../../utils/localStorage';
import { Injectable } from '@angular/core';
import {
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class HttpTokenInterceptor implements HttpInterceptor {
  constructor() {}

  intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    let headersConfig: any = {};

    if (!req.url.includes('/api/file')) {
      headersConfig = {
        'Content-Type': 'application/json',
        Accept: 'application/json',
      };
    }

    if (
      req.url.includes(environment.ApiUrl) ||
      req.url.includes(environment.FileServer)
    ) {
      const token = AppStorage.getItem(StorageKeys.Token);

      if (token) {
        headersConfig['Authorization'] = `${token}`;
      }
    }

    const request = req.clone({ setHeaders: headersConfig });
    return next.handle(request);
  }
}
