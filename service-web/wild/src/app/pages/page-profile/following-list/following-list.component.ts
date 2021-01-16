import { Component, OnInit } from '@angular/core';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';

@Component({
  selector: 'app-following-list',
  templateUrl: './following-list.component.html',
  styleUrls: ['./following-list.component.scss'],
})
export class FollowingListComponent extends BaseComponent implements OnInit {
  constructor(private authenticationService: AuthenticationService) {
    super();
  }
  Members: AppModels.MemberModel[] = [];

  ngOnInit(): void {
    this.load();
  }

  load() {
    this.authenticationService.Followings({}).subscribe(
      (v) => {
        this.Members = v.Members;
      },
      (err) => {
        console.log('err:', err);
      }
    );
  }

  unfollow(memberID: string | undefined) {
    this.authenticationService
      .Follow({
        Follow: false,
        MemberID: memberID ?? '',
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
