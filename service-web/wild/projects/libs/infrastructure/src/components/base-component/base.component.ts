import { Component } from '@angular/core';
import { Constants } from '../../utils/Globals';
import { ValidationErrTypes } from '../../utils/enumerations';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'lib-base',
  templateUrl: './base.component.html',
  styleUrls: ['./base.component.scss'],
})
export class BaseComponent {
  Constants = Constants;
  ValidationErrTypes = ValidationErrTypes;

  AddBaseUrlFileServer(url: string | undefined) {
    //http://192.168.1.101/api/file/download/f-186184211.jpg
    return environment.FileServer + 'file/download/' + url;
  }

  ColorClass(Color: string, more: any = {}) {
    switch (Color) {
      case '#187B3B':
        return { color1: true, ...more };
      case '#CB6941':
        return { color2: true, ...more };
      case '#000000':
        return { color3: true, ...more };
      case '#7AB6C1':
        return { color4: true, ...more };
    }
  }
}
