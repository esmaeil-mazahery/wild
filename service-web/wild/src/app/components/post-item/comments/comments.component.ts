import { Component, Input, OnInit } from '@angular/core';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';

@Component({
  selector: 'app-comments',
  templateUrl: './comments.component.html',
  styleUrls: ['./comments.component.scss'],
})
export class CommentsComponent implements OnInit {
  @Input() list?: AppModels.CommentModel[];

  constructor() {}

  ngOnInit(): void {}
}
