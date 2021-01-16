import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AlertService } from 'projects/libs/infrastructure/src/serivces/system/alert.service';

@Component({
  templateUrl: './page-signup.component.html',
  styleUrls: ['./page-signup.component.scss'],
})
export class PageSignupComponent implements OnInit {
  hide = true;
  returnUrl!: string;
  form!: FormGroup;
  loading = false;
  constructor(
    private formBuilder: FormBuilder,
    private route: ActivatedRoute,
    private router: Router,
    private alertService: AlertService,
    private authService: AuthenticationService
  ) {}

  ngOnInit(): void {
    this.returnUrl = this.route.snapshot.queryParams.returnUrl || '/';

    this.form = this.formBuilder.group({
      Username: ['', [Validators.required]],
      Name: ['', [Validators.required]],
      Family: ['', [Validators.required]],
      Mobile: ['', [Validators.required]],
      Email: ['', [Validators.required]],
      Password: ['', [Validators.required]],
    });
  }

  get f() {
    return this.form.controls;
  }

  onSubmit() {
    this.loading = true;
    this.authService
      .Register({
        Member: {
          Username: this.f.Username.value,
          Name: this.f.Name.value,
          Family: this.f.Family.value,
          Mobile: this.f.Mobile.value,
          Email: this.f.Email.value,
          Password: this.f.Password.value,
        },
      })
      .subscribe(
        (v) => {
          this.loading = false;
          this.router.navigateByUrl(this.returnUrl);
        },
        (err) => {
          if (err.error.code == 6) {
            this.alertService.openSnackBar(
              'این نام کاربری در سیستم وجود دارد',
              false,
              3000
            );
          } else {
            this.alertService.openSnackBar('خطایی رخ داد', false, 3000);
          }
          this.loading = false;
        }
      );
  }
}
