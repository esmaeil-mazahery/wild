export class Guid {
  public static newGuid() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(
      /[xy]/g,
      function (c) {
        var r = (Math.random() * 16) | 0,
          v = c == 'x' ? r : (r & 0x3) | 0x8;
        return v.toString(16);
      }
    );
  }
}

export class Extention {
  public static IsEmpty(val: any) {
    return val == null || val == '' || val == undefined;
  }

  public static ConvertString(val: any): string {
    return Extention.IsEmpty(val) ? '' : val;
  }
}
