import { Injectable } from "@angular/core";
import { Router, NavigationStart } from "@angular/router";
import { Observable, Subject } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class AlertService {
  private subject = new Subject<any>();
  private keepAfterNavigationChange = false;

  constructor(private router: Router) {
    router.events.subscribe(event => {
      if (event instanceof NavigationStart) {
        if (this.keepAfterNavigationChange) {
          this.keepAfterNavigationChange = false;
        } else {
          // clear alert
          this.subject.next();
        }
      }
    });
  }

  success(message: string, keepAfterNavigationChange = false) {
    this.keepAfterNavigationChange = keepAfterNavigationChange;
    this.subject.next({ type: "success", text: message,error:false });
  }

  error(message: string, keepAfterNavigationChange = false) {
    this.keepAfterNavigationChange = keepAfterNavigationChange;
    this.subject.next({ type: "error", text: message ,error:true});
  }

  getMessage(): Observable<any> {
    return this.subject.asObservable();
  }

  openSnackBar(message: string, error: boolean, duration: number) {
    this.subject.next({
      type: "SnackBar",
      text: message,
      error: error,
      data: {
        duration: duration
      }
    });
  }

  openDialog(message: string, error: boolean, template: any) {
    this.subject.next({
      type: "Dialog",
      text: message,
      error:error,
      data: {
        template: template
      }
    });
  }

  showText(message: string, error: boolean) {
    this.subject.next({
      type: "Text",
      text: message,
      error:error
    });
  }
}
