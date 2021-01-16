import { Component, OnInit, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import {
  AuthenticationModels,
  AuthenticationService,
} from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';
import { UploadService } from 'projects/libs/infrastructure/src/serivces/Entity/upload.service';
import { AlertService } from 'projects/libs/infrastructure/src/serivces/system/alert.service';
import { forkJoin, Observable } from 'rxjs';

@Component({
  selector: 'app-desktop-menu',
  templateUrl: './desktop-menu.component.html',
  styleUrls: ['./desktop-menu.component.scss'],
})
export class DesktopMenuComponent extends BaseComponent implements OnInit {
  profile!: AppModels.MemberModel;

  constructor(
    public uploadService: UploadService,
    private authService: AuthenticationService,
    private alertService: AlertService,
    private router: Router
  ) {
    super();
  }

  ngOnInit(): void {
    this.authService.GetProfileSync().subscribe((p) => {
      this.profile = p;
    });
  }

  LogOut() {
    this.authService.LogOut();
    this.router.navigateByUrl('/auth');
  }

  inLocation(location: string) {
    return this.router.url === location;
  }
}
