import { PageSigninComponent } from './pages/page-signin/page-signin.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuardService } from 'projects/libs/infrastructure/src/serivces/system/auth-guard.service';
import { PageHomeComponent } from './pages/page-home/page-home.component';
import { PageSignupComponent } from './pages/page-signup/page-signup.component';
import { PageProfileComponent } from './pages/page-profile/page-profile.component';
import { PageSearchComponent } from './pages/page-search/page-search.component';
import { PageNotificationComponent } from './pages/page-notification/page-notification.component';

const routes: Routes = [
  {
    path: '',
    component: PageHomeComponent,
    canActivate: [AuthGuardService],
  },
  {
    path: 'profile',
    component: PageProfileComponent,
    canActivate: [AuthGuardService],
  },
  {
    path: 'search',
    component: PageSearchComponent,
    canActivate: [AuthGuardService],
  },
  {
    path: 'notification',
    component: PageNotificationComponent,
    canActivate: [AuthGuardService],
  },
  {
    path: 'auth',
    component: PageSigninComponent,
  },
  {
    path: 'signup',
    component: PageSignupComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
