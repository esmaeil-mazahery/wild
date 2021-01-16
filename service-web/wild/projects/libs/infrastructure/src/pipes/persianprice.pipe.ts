import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'persianprice',
})
export class PersianPricePipe implements PipeTransform {
  transform(value?: number, args?: any): any {
    if (value != null && value > 0) {
      return value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
    } else {
      return '';
    }
  }
}
