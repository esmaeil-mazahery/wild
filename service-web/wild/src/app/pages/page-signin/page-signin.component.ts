import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { AlertService } from 'projects/libs/infrastructure/src/serivces/system/alert.service';

@Component({
  templateUrl: './page-signin.component.html',
  styleUrls: ['./page-signin.component.scss'],
})
export class PageSigninComponent implements OnInit {
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
      Password: ['', [Validators.required]],
    });
  }

  get f() {
    return this.form.controls;
  }

  onSubmit() {
    this.loading = true;
    this.authService
      .Login({
        Username: this.f.Username.value,
        Password: this.f.Password.value,
      })
      .subscribe(
        (v) => {
          this.loading = false;
          this.router.navigateByUrl(this.returnUrl);
        },
        (err) => {
          this.alertService.openSnackBar(
            'نام کاربری یا رمز عبور اشتباه است',
            false,
            3000
          );

          this.loading = false;
        }
      );
  }
}
