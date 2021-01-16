import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { PostService } from 'projects/libs/infrastructure/src/serivces/Entity/post.service';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';
import { Component, OnInit } from '@angular/core';

@Component({
  templateUrl: './page-home.component.html',
  styleUrls: ['./page-home.component.scss'],
})
export class PageHomeComponent implements OnInit {
  constructor(
    private postService: PostService,
    private authenticationService: AuthenticationService
  ) {}

  currentPage: number = 1;
  ngOnInit(): void {
    this.loadPosts();
  }

  postList: AppModels.PostModel[] = [];

  loadPosts() {
    this.postService
      .List({
        Page: this.currentPage,
      })
      .subscribe(
        (v) => {
          this.postList = v.Posts;
        },
        (err) => {
          console.log('err:', err);
        }
      );
  }

  addedPost(post: AppModels.PostModel) {
    this.postList = [
      {
        ...post,
        Member: this.authenticationService.GetProfile(),
      },
      ...this.postList,
    ];
  }
}
