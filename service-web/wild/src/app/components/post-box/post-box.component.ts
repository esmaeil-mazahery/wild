import { PostService } from 'projects/libs/infrastructure/src/serivces/Entity/post.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-post-box',
  templateUrl: './post-box.component.html',
  styleUrls: ['./post-box.component.scss'],
})
export class PostBoxComponent implements OnInit {
  constructor(private postService: PostService) {}

  ngOnInit(): void {}
}
