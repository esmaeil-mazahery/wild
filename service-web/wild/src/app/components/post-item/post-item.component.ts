import { Input } from '@angular/core';
import { Component, OnInit } from '@angular/core';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';
import { PostService } from 'projects/libs/infrastructure/src/serivces/Entity/post.service';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { CommentService } from 'projects/libs/infrastructure/src/serivces/Entity/comment.service';
import { AlertService } from 'projects/libs/infrastructure/src/serivces/system/alert.service';

@Component({
  selector: 'app-post-item',
  templateUrl: './post-item.component.html',
  styleUrls: ['./post-item.component.scss'],
})
export class PostItemComponent extends BaseComponent implements OnInit {
  @Input() post!: AppModels.PostModel;
  showComments: boolean = false;
  loading: boolean = false;
  constructor(
    private postService: PostService,
    private commentService: CommentService,
    private alertService: AlertService,
    private authenticationService: AuthenticationService
  ) {
    super();
  }

  ngOnInit(): void {}

  like() {
    if (this.loading == false) {
      this.loading = true;
      this.postService
        .Like({
          ID: this.post.ID ?? '',
          Like: !this.post.MemberLike,
        })
        .subscribe(
          (v) => {
            this.post.MemberLike = v.Result;
            this.loading = false;
          },
          (err) => {
            this.loading = false;
            this.post.MemberLike = !this.post.MemberLike;
            this.alertService.openSnackBar(
              'خطا در برقراری ارتباط با سرور',
              true,
              5000
            );
          }
        );

      this.post.MemberLike = !this.post.MemberLike;
    }
  }
  comments?: AppModels.CommentModel[];
  currentCommentPage: number = 1;
  comment() {
    if (!this.showComments && this.comments == null) {
      this.commentService
        .List({
          Page: this.currentCommentPage,
          PostID: this.post.ID,
        })
        .subscribe(
          (v) => {
            this.showComments = !this.showComments;
            this.comments = v.Comments;
          },
          (err) => {
            console.log('err:', err);
          }
        );
    } else {
      this.showComments = !this.showComments;
    }
  }

  onAddedComment(comment: AppModels.CommentModel) {
    this.comments = [
      {
        ...comment,
        Member: this.authenticationService.GetProfile(),
      },
      ...(this.comments ?? []),
    ];
  }
}
