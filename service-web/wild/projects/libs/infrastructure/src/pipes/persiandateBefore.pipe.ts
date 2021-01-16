import { Pipe, PipeTransform } from '@angular/core';
import * as moment from 'jalali-moment';

@Pipe({
  name: 'persiandatebefore'
})
export class PersiandateBeforePipe implements PipeTransform {

  transform(value?: any, args?: any): any {
    if (value && value != "0001-01-01T00:00:00") {
      let MomentDate = moment(value, 'YYYY/MM/DD');
      const x = moment().diff(MomentDate, 'days');
      switch (true) {
        case 0 === x:
          return "امروز";
        case 1 === x:
          return "یک روز پیش";
        case 2 === x:
          return "دو روز پیش";
        case 3 === x:
          return "سه روز پیش";
        case 4 === x:
          return "چهار روز پیش";
        case 5 === x:
          return "پنج روز پیش";
        case 6 === x:
          return "شش روز پیش";
        case x < 14:
          return "یک هفته پیش";
        case x < 21:
          return "دو هفته پیش";
        case x < 28:
          return "سه هفته پیش";
        case x < 60:
          return "یک ماه پیش";
        case x < 90:
          return "دو ماه پیش";
        case x < 120:
          return "سه ماه پیش";
        default:
          return MomentDate.locale('fa').format('YYYY/M/D');
      }
    } else {
      return "";
    }
  }
}
