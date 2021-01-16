import { HttpParams } from '@angular/common/http';

export function BuildHttpParamsWithModel(model: any): HttpParams {
  let httpParams = new HttpParams();
  Object.keys(model).forEach((key) => {
    const value = model[key];

    if (value instanceof Array) {
      // httpParams = httpParams.append(key, value.join(','));
      value.forEach((item) => {
        if (item instanceof Array || item instanceof Object) {
          httpParams = httpParams.append(key, JSON.stringify(item));
        } else if (item !== null && item !== undefined) {
          httpParams = httpParams.append(key, item);
        }
      });
    } else if (value instanceof Object) {
      httpParams = httpParams.append(key, JSON.stringify(value));
    } else if (value !== null && value !== undefined) {
      httpParams = httpParams.append(key, value);
    }
  });

  return httpParams;
}
