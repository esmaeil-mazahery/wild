import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { PageSigninComponent } from './pages/page-signin/page-signin.component';
import { PageSignupComponent } from './pages/page-signup/page-signup.component';
import { PageHomeComponent } from './pages/page-home/page-home.component';
import { InfrastructureModule } from 'projects/libs/infrastructure/src/public-api';
import { PostListComponent } from './components/post-list/post-list.component';
import { DesktopMenuComponent } from './components/desktop-menu/desktop-menu.component';
import { MobileMenuComponent } from './components/mobile-menu/mobile-menu.component';
import { PostBoxComponent } from './components/post-box/post-box.component';
import { PostCreateComponent } from './components/post-create/post-create.component';
import { PageProfileComponent } from './pages/page-profile/page-profile.component';
import { PageSearchComponent } from './pages/page-search/page-search.component';
import { PageNotificationComponent } from './pages/page-notification/page-notification.component';
import { FollowerListComponent } from './pages/page-profile/follower-list/follower-list.component';
import { FollowingListComponent } from './pages/page-profile/following-list/following-list.component';
import { PostItemComponent } from './components/post-item/post-item.component';
import { CommentsComponent } from './components/post-item/comments/comments.component';
import { SearchBoxComponent } from './components/search-box/search-box.component';
import { CommentCreateComponent } from './components/post-item/comments/comment-create/comment-create.component';
import { CommentItemComponent } from './components/post-item/comments/comment-item/comment-item.component';
import { SuggestionComponent } from './components/suggestion/suggestion.component';
import { NotifyListComponent } from './pages/page-notification/notify-list/notify-list.component';

@NgModule({
  declarations: [
    AppComponent,
    PageSigninComponent,
    PageSignupComponent,
    PageHomeComponent,
    PostListComponent,
    DesktopMenuComponent,
    MobileMenuComponent,
    PostBoxComponent,
    PostCreateComponent,
    PageProfileComponent,
    PageSearchComponent,
    PageNotificationComponent,
    FollowerListComponent,
    FollowingListComponent,
    PostItemComponent,
    CommentsComponent,
    SearchBoxComponent,
    CommentCreateComponent,
    CommentItemComponent,
    SuggestionComponent,
    NotifyListComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,

    InfrastructureModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
