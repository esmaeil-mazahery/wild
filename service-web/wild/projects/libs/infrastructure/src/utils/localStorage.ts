import { Extention } from './generic';

export enum StorageKeys {
  Token = 1,
  Mobile = 2,
  Username = 3,
  Email = 4,
  Image = 5,
  Name = 6,
  Family = 7,
  ImageHeader = 8,
  Follower = 9,
  Following = 10,
  Biography = 11,
}

export class AppStorage {
  public static CreateUUID() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (c) => {
      const r = (Math.random() * 16) | 0,
        v = c === 'x' ? r : (r & 0x3) | 0x8;
      return v.toString(16);
    });
  }

  public static SetNotNull = (key: StorageKeys, value: any) => {
    if (!Extention.IsEmpty(value)) {
      AppStorage.setItem(key, value);
    }
  };

  public static setItem(key: StorageKeys, value: any, type?: any): void {
    if (
      value instanceof String ||
      typeof value === 'string' ||
      type === String
    ) {
      localStorage.setItem(key.toString(), Extention.ConvertString(value));
    } else {
      localStorage.setItem(key.toString(), JSON.stringify(value));
    }
  }

  public static getItem(key: StorageKeys, type?: any) {
    const valueString = localStorage.getItem(key.toString());

    if (valueString) {
      return type === Number
        ? Number.parseFloat(valueString)
        : type === Array
        ? valueString.split(',')
        : type === Object
        ? JSON.parse(valueString)
        : type === Boolean
        ? ConvertBoolean(valueString)
        : valueString;
    } else {
      return type === Number
        ? 0
        : type === Array
        ? null
        : type === Object
        ? null
        : type === Boolean
        ? false
        : '';
    }
  }

  public static removeItem(key: StorageKeys) {
    localStorage.removeItem(key.toString());
  }

  public static clear() {
    localStorage.clear();
  }
}

export const ConvertBoolean = (value: any) => {
  return value === 'false' ? false : value;
};
