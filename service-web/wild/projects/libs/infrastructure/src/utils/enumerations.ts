export enum ValidationErrTypes {
  minlength,
  maxlength,
  required,
  pending,
  taken,
  persian,
  regex,
}

export class EnumEx {
  private constructor() {}

  static getDescriptionsAndValues<T extends number>(e: any) {
    return EnumEx.getNames(e).map((n) => ({
      description: e[n + 'Description'],
      value: e[n] as T,
    }));
  }

  static getNamesAndValues<T extends number>(e: any) {
    return EnumEx.getNames(e).map((n) => ({ name: n, value: e[n] as T }));
  }

  static getNames(e: any) {
    return Object.keys(e).filter((k) => typeof e[k] === 'number') as string[];
  }

  static getValues<T extends number>(e: any) {
    return Object.keys(e)
      .map((k) => e[k])
      .filter((v) => typeof v === 'number') as T[];
  }
}

export enum EventBusActions {
  Add = 1,
  Delete = 2,
  Edit = 3,
  DeleteAll = 4,
  AddCollection = 5,
  Select = 6,
  DeSelect = 7,
  DeSelectAll = 8,
  changeRegisterItem = 9,
  EmitRegister = 10,
  hide = 11,
  show = 12,
  toggle = 13,
}

export enum ProductType {
  Food = 1,
  Fruit = 2,
  WasherAndCleaner = 3,
}

export enum UnitType {
  Number = 1,
  NumberDescription = 'تعداد',
  Kg = 2,
  KgDescription = 'کیلوگرم',
  Liter = 3,
  LiterDescription = 'لیتر',
  Mesghal = 4,
  MesghalDescription = 'مثقال',
  Gram = 5,
  GramDescription = 'گرم',
  Metr = 6,
  MetrDescription = 'متر',
}

export enum UnitOfWorkTypes {
  Resseller = 1,
  CentralOffice = 2,
}

export enum SocialNetworkType {
  FaceBook = 1,
  Twitter = 2,
  Telegram = 3,
  Instagram = 4,
  Linkedin = 5,
  googleplus = 6,
}
