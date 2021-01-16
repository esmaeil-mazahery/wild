import { NgModule } from '@angular/core';
import { BaseComponent } from '../components/base-component/base.component';
import { SnakbarAlarmComponent } from '../components/snakbarAlarm/snakbarAlarm.component';
import { PersiandatePipe } from '../pipes/persiandate.pipe';
import { PersiandateBeforePipe } from '../pipes/persiandateBefore.pipe';
import { PersianPricePipe } from '../pipes/persianprice.pipe';
import { SharedModule } from './shared.module';

const components = [BaseComponent, SnakbarAlarmComponent];

const icons: any[] = [];

const pipes = [PersiandatePipe, PersianPricePipe, PersiandateBeforePipe];

const validator: any[] = [];

const directives: any[] = [];

@NgModule({
  declarations: [
    ...components,
    ...icons,
    ...pipes,
    ...validator,
    ...directives,
  ],
  imports: [SharedModule],
  exports: [
    SharedModule,
    ...components,
    ...icons,
    ...pipes,
    ...validator,
    ...directives,
  ],
})
export class InfrastructureModule {}
