import { BehaviorSubject } from "rxjs";
import { Injectable } from "@angular/core";

@Injectable({
  providedIn: "root",
})
export class MessagingService {
  Messages = new BehaviorSubject(null);
  constructor() {}

  public next(value) {
    this.Messages.next(value);
  }

  public getSubject() {
    return this.Messages;
  }
}
