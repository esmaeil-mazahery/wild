import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MainBusService {
  private publicListSource = new Subject<any>();
  publicListList$ = this.publicListSource.asObservable();

  publicListSourceEmit(data: any, type: string, trigger: any) {
    this.publicListSource.next({ data: data, type: type, trigger: trigger });
  }
  constructor() { }

}
