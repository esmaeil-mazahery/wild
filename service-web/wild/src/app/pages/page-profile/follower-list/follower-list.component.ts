import { Component, OnInit } from '@angular/core';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';

@Component({
  selector: 'app-follower-list',
  templateUrl: './follower-list.component.html',
  styleUrls: ['./follower-list.component.scss'],
})
export class FollowerListComponent extends BaseComponent implements OnInit {
  constructor(private authenticationService: AuthenticationService) {
    super();
  }
  Members: AppModels.MemberModel[] = [];

  ngOnInit(): void {
    this.load();
  }

  load() {
    this.authenticationService.Followers({}).subscribe(
      (v) => {
        this.Members = v.Members;
      },
      (err) => {
        console.log('err:', err);
      }
    );
  }

  follow(memberID: string | undefined) {
    this.authenticationService
      .Follow({
        Follow: true,
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
