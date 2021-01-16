import { environment } from './../environments/environment';
import { Constants } from './Globals';
import { startWith } from "rxjs/operators";
import { ValidationErrTypes } from './enumerations';

export class GlobalsBase {
  Constants = Constants;
  ValidationErrTypes = ValidationErrTypes;

  AddBaseUrlFileServer(url) {
    return environment.FileServer + "Files/"+url;
  }
}
