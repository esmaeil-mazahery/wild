import { Injectable } from '@angular/core';
import {
  HttpClient,
  HttpRequest,
  HttpEventType,
  HttpResponse,
  HttpHeaders,
} from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { map, catchError } from 'rxjs/operators';
import { AlertService } from '../system/alert.service';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root',
})
export class UploadService {
  siteUrl = `${environment.FileServer}file`;

  constructor(private http: HttpClient) {}

  public upload(
    files: Set<File>
  ): { [key: string]: { progress: Observable<number>; filename: string } } {
    // this will be the our resulting map
    const status: {
      [key: string]: { progress: Observable<number>; filename: string };
    } = {};

    files.forEach((file) => {
      // create a new multipart-form for every file
      const formData: FormData = new FormData();
      formData.append('file', file, file.name);

      // create a http-post request and pass the form
      // tell it to report the upload progress
      const req = new HttpRequest('post', this.siteUrl, formData, {
        reportProgress: true,
      });

      // create a new progress-subject for every file
      const progress = new Subject<number>();
      let Url = '';
      // send the http-request and subscribe for progress-updates
      this.http.request<any>(req).subscribe((event) => {
        if (event.type === HttpEventType.UploadProgress) {
          // calculate the progress percentage
          const percentDone = Math.round((100 * event.loaded) / event.total!);

          // pass the percentage into the progress-stream
          progress.next(percentDone);
        } else if (event instanceof HttpResponse) {
          // Close the progress-stream if we get an answer form the API
          // The upload is complete

          Url = event.body['Url'];
          status[file.name] = {
            progress: progress.asObservable(),
            filename: Url,
          };

          progress.complete();
        }
      });

      // Save every progress-observable in a map of all observables
      status[file.name] = {
        progress: progress.asObservable(),
        filename: Url,
      };
    });

    // return the map of progress.observables
    return status;
  }

  public exist(Filename: string) {}

  public delete(Filename: string) {
    return this.http.delete(this.siteUrl + '/delete').pipe(map((a) => true));
  }
}
