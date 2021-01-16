import { Component, OnInit } from '@angular/core';
import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';
import { PostService } from 'projects/libs/infrastructure/src/serivces/Entity/post.service';

@Component({
  templateUrl: './page-search.component.html',
  styleUrls: ['./page-search.component.scss'],
})
export class PageSearchComponent implements OnInit {
  currentPage: number = 1;
  postList: AppModels.PostModel[] = [];
  constructor(
    private postService: PostService,
    private authenticationService: AuthenticationService
  ) {}

  ngOnInit(): void {}

  search(term: string) {
    this.postService
      .Search({
        Page: this.currentPage,
        Term: term,
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
}
