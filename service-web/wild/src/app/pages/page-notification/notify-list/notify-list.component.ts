import { NotifyService } from 'projects/libs/infrastructure/src/serivces/Entity/notify.service';
import { Component, OnInit } from '@angular/core';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';

@Component({
  selector: 'app-notify-list',
  templateUrl: './notify-list.component.html',
  styleUrls: ['./notify-list.component.scss'],
})
export class NotifyListComponent extends BaseComponent implements OnInit {
  constructor(private notifyService: NotifyService) {
    super();
  }
  page: number = 1;
  Notifies: AppModels.NotifyModel[] = [];

  NotifyType = AppModels.NotifyType;

  ngOnInit(): void {
    this.load();
  }

  load() {
    this.notifyService
      .List({
        Page: this.page,
      })
      .subscribe(
        (v) => {
          this.Notifies = v.Notifies;
          this.read();
        },
        (err) => {
          console.log('err:', err);
        }
      );
  }

  read() {
    this.notifyService
      .Read({
        IDs: this.Notifies.map((m) => m.ID ?? ''),
      })
      .subscribe(
        (v) => {
          this.load();
        },
        (err) => {
          console.log('err:', err);
        }
      );
  }
}
