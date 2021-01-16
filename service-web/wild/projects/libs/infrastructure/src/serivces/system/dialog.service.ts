import { Subject, Observable } from 'rxjs';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class DialogService {
  private subject = new Subject<any>();

  constructor() { }

  getMessage(): Observable<any> {
    return this.subject.asObservable();
  }

  next(data: any) {
    this.subject.next(data);
  }
}
