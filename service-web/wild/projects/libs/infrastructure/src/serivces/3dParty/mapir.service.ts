import { Injectable } from '@angular/core';
import { HttpHeaders, HttpClient, HttpParams } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class MapirService {

  baseurl = "https://map.ir"
  constructor(private http: HttpClient) {
  }

  autocomplete(data) {
    return this.http.post<any>(this.baseurl + "/search/autocomplete", data, {
      headers: new HttpHeaders({
        "Content-Type": "application/json",
        "x-api-key": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjBjM2IxOWVhNjRlOWQ4NjdiZWJiYmQ0NTI0MzUyZGQ3ODYwOGQyNDgyZTY4ODM0NTRjNzFmNGRmNjMzYTI0YmFiMjcwMjM5MThkYWY0NDE0In0.eyJhdWQiOiI3MTk0IiwianRpIjoiMGMzYjE5ZWE2NGU5ZDg2N2JlYmJiZDQ1MjQzNTJkZDc4NjA4ZDI0ODJlNjg4MzQ1NGM3MWY0ZGY2MzNhMjRiYWIyNzAyMzkxOGRhZjQ0MTQiLCJpYXQiOjE1NzcwNDk4MDUsIm5iZiI6MTU3NzA0OTgwNSwiZXhwIjoxNTc5NTU1NDA1LCJzdWIiOiIiLCJzY29wZXMiOlsiYmFzaWMiXX0.p1gM6nh1CzFM2VA6hjFfJOnThSU96_7OLdo-YK3Op_BtGaM5tuVno-4W2bx7rr05utgJ7Dj-BvyAT8-NxDhrG0q0q_9dgzZBchARlRFS3SGlTmLqn7-Mcrix1Eu9hAjVEJ3epdpH6lrHMlE59QKOSXTT6bpavFNhCZWaq1--ysFiGppO0y7SFFdz9TlV0MHdOsSnIBpdmgOx47R7SHkHEx3aYBm2MQ6qHWz_FfXjuk6R0EStC9LPhIsGFOgxAhxCd1-xAiZ7BbW_K7INYXN_KcsQXZ0ff58Z9fxANt5T6qI_60s_cTsTbEW2PVmGLMwQNX8UpuQOnTzsNcO8DFVvTw",
      })
    });
  }


  reverse(lat, lng) {
    let httpParams = new HttpParams();
    httpParams = httpParams.append("lat", lat.toString());
    httpParams = httpParams.append("lon", lng.toString());
    return this.http.get<any>(this.baseurl + "/reverse", {
      headers: new HttpHeaders({
        "Content-Type": "application/json",
        "x-api-key": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjBjM2IxOWVhNjRlOWQ4NjdiZWJiYmQ0NTI0MzUyZGQ3ODYwOGQyNDgyZTY4ODM0NTRjNzFmNGRmNjMzYTI0YmFiMjcwMjM5MThkYWY0NDE0In0.eyJhdWQiOiI3MTk0IiwianRpIjoiMGMzYjE5ZWE2NGU5ZDg2N2JlYmJiZDQ1MjQzNTJkZDc4NjA4ZDI0ODJlNjg4MzQ1NGM3MWY0ZGY2MzNhMjRiYWIyNzAyMzkxOGRhZjQ0MTQiLCJpYXQiOjE1NzcwNDk4MDUsIm5iZiI6MTU3NzA0OTgwNSwiZXhwIjoxNTc5NTU1NDA1LCJzdWIiOiIiLCJzY29wZXMiOlsiYmFzaWMiXX0.p1gM6nh1CzFM2VA6hjFfJOnThSU96_7OLdo-YK3Op_BtGaM5tuVno-4W2bx7rr05utgJ7Dj-BvyAT8-NxDhrG0q0q_9dgzZBchARlRFS3SGlTmLqn7-Mcrix1Eu9hAjVEJ3epdpH6lrHMlE59QKOSXTT6bpavFNhCZWaq1--ysFiGppO0y7SFFdz9TlV0MHdOsSnIBpdmgOx47R7SHkHEx3aYBm2MQ6qHWz_FfXjuk6R0EStC9LPhIsGFOgxAhxCd1-xAiZ7BbW_K7INYXN_KcsQXZ0ff58Z9fxANt5T6qI_60s_cTsTbEW2PVmGLMwQNX8UpuQOnTzsNcO8DFVvTw"
      }),
      params: httpParams
    })
  }
}
