import { Injectable } from '@angular/core';
// import * as L from "leaflet";

@Injectable()
export class Globals {
  public static ApiUrl: string = 'https://localhost:44365/api'; //environment.ApiUrl;//
  public static FileServer: string = 'https://localhost:44305/'; // environment.FileServer;
}

@Injectable()
export class Constants {
  public static PatternMobile: RegExp = /^[0]{0,1}[9][0-9]{9}$/i;
  public static PatternPhone: RegExp = /^[0]{0,1}[1-9][0-9]{9}$/i;
  public static PatternMobileOrEmail: RegExp = /^([0]{0,1}[9][0-9]{9,9})|(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$/i;
  public static PatternEmail: RegExp = /^(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$/i;
  public static PatternWebsite: RegExp = /^(http\:\/\/|https\:\/\/|HTTPS\:\/\/|HTTP\:\/\/)?([A-Za-z0-9][A-Za-z0-9\-]*\.)+[A-Za-z0-9][A-Za-z0-9\-]*$/i;

  public static PasswordMaxLength = 50;
  public static NameMaxLength = 55;
  public static NameMinLength = 2;
  public static TitleLength = 250;
  public static Description = 4000;

  public static LimitActivities = 5;
  public static LimitBusinesses = 5;
  public static LimitActivitiesPerBusiness = 5;

  public static GoogleReCaptchaKey = '6LftW9EZAAAAALsbN8kPsnVoh_FrGRSghu_zougr';
}
