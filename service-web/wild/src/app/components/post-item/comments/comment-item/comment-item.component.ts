import { BaseComponent } from './../../../../../../projects/libs/infrastructure/src/components/base-component/base.component';
import { Component, Input, OnInit } from '@angular/core';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';
import { CommentService } from 'projects/libs/infrastructure/src/serivces/Entity/comment.service';
import { AlertService } from 'projects/libs/infrastructure/src/serivces/system/alert.service';

@Component({
  selector: 'app-comment-item',
  templateUrl: './comment-item.component.html',
  styleUrls: ['./comment-item.component.scss'],
})
export class CommentItemComponent extends BaseComponent implements OnInit {
  @Input() comment!: AppModels.CommentModel;
  loading: boolean = false;
  constructor(
    private commentService: CommentService,
    private alertService: AlertService
  ) {
    super();
  }

  ngOnInit(): void {}

  like() {
    if (this.loading == false) {
      this.loading = true;
      this.commentService
        .Like({
          ID: this.comment.ID ?? '',
          Like: !this.comment.MemberLike,
        })
        .subscribe(
          (v) => {
            this.comment.MemberLike = v.Result;
            this.loading = false;
          },
          (err) => {
            this.loading = false;
            this.comment.MemberLike = !this.comment.MemberLike;
            this.alertService.openSnackBar(
              'خطا در برقراری ارتباط با سرور',
              true,
              5000
            );
          }
        );

      this.comment.MemberLike = !this.comment.MemberLike;
    }
  }
}
