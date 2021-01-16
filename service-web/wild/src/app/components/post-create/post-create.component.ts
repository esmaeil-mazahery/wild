import { UploadService } from 'projects/libs/infrastructure/src/serivces/Entity/upload.service';
import {
  AuthenticationModels,
  AuthenticationService,
} from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { Component, EventEmitter, OnInit, ViewChild } from '@angular/core';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import { FormControl } from '@angular/forms';
import {
  PostModels,
  PostService,
} from 'projects/libs/infrastructure/src/serivces/Entity/post.service';
import { forkJoin, Observable } from 'rxjs';
import { Output } from '@angular/core';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';

@Component({
  selector: 'app-post-create',
  templateUrl: './post-create.component.html',
  styleUrls: ['./post-create.component.scss'],
})
export class PostCreateComponent extends BaseComponent implements OnInit {
  @Output() added = new EventEmitter<AppModels.PostModel>();
  profile!: AppModels.MemberModel;
  Content = new FormControl();
  constructor(
    private authService: AuthenticationService,
    private postService: PostService,
    private uploadService: UploadService
  ) {
    super();
  }

  @ViewChild('file') file: any;
  uploading = false;
  uploadSuccessful = false;
  public files: Set<File> = new Set();
  image: string | undefined;
  progress:
    | {
        [key: string]: { progress: Observable<number>; filename: string };
      }
    | undefined;
  progressNumber: number = 0;

  addFiles() {
    this.file.nativeElement.click();
  }

  onFilesAdded() {
    const files: { [key: string]: File } = this.file.nativeElement.files;
    for (const key in files) {
      if (!isNaN(parseInt(key, 10))) {
        this.uploadSuccessful = false;
        this.files.clear();
        this.files.add(files[key]);
      }
    }
    this.upload();
  }

  upload() {
    // set the component state to "uploading"
    this.uploading = true;
    // start the upload and save the progress map
    this.progress = this.uploadService.upload(this.files);

    // convert the progress map into an array
    const allProgressObservables = [];
    for (const key in this.progress) {
      if (this.progress.hasOwnProperty(key)) {
        allProgressObservables.push(this.progress[key].progress);

        this.progress[key].progress.subscribe((p) => {
          this.progressNumber = p;
        });
      }
    }

    // When all progress-observables are completed...
    forkJoin(allProgressObservables).subscribe((_) => {
      for (const key in this.progress) {
        if (this.progress.hasOwnProperty(key)) {
          this.image = this.progress[key].filename;
          console.log(this.image);
        }
      }

      // ... the upload was successful...
      this.uploadSuccessful = true;

      // ... and the component is no longer uploading
      this.uploading = false;
      this.progressNumber = 0;
    });
  }

  ngOnInit(): void {
    this.authService.GetProfileSync().subscribe((p) => {
      this.profile = p;
    });
  }

  onSubmit() {
    this.postService
      .Register({
        Post: {
          Image: this.image,
          Content: this.Content.value,
          Likes: [],
          Comments: [],
        },
      })
      .subscribe(
        (v) => {
          this.added.emit({
            ID: v.ID,
            Content: this.Content.value,
            Image: this.image,
            Likes: [],
            Comments: [],
          });

          this.Content.setValue('');
          this.image = '';
          this.uploadSuccessful = false;
        },
        (err) => {}
      );
  }
}
