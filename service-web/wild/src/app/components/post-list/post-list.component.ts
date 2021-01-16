import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';
import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-post-list',
  templateUrl: './post-list.component.html',
  styleUrls: ['./post-list.component.scss'],
})
export class PostListComponent implements OnInit {
  @Input() list: AppModels.PostModel[] = [];

  constructor() {}

  ngOnInit(): void {}
}
