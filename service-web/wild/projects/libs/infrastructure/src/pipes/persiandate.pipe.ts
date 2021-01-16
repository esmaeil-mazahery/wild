import { Pipe, PipeTransform } from "@angular/core";
import * as moment from "jalali-moment";

@Pipe({
  name: "persiandate",
})
export class PersiandatePipe implements PipeTransform {
  transform(value?: any, args?: any): any {
    if (value && value != "0001-01-01T00:00:00") {
      let MomentDate = moment(value);
      return MomentDate.locale("fa").format("YYYY/M/D");
    } else {
      return "";
    }
  }
}
