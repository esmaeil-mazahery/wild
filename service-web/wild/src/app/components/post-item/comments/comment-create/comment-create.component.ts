import {
  Component,
  EventEmitter,
  Input,
  OnInit,
  Output,
  ViewChild,
} from '@angular/core';
import { FormControl } from '@angular/forms';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import { AuthenticationService } from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { CommentService } from 'projects/libs/infrastructure/src/serivces/Entity/comment.service';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';
import { UploadService } from 'projects/libs/infrastructure/src/serivces/Entity/upload.service';
import { forkJoin, Observable } from 'rxjs';

@Component({
  selector: 'app-comment-create',
  templateUrl: './comment-create.component.html',
  styleUrls: ['./comment-create.component.scss'],
})
export class CommentCreateComponent extends BaseComponent implements OnInit {
  @Input() postID?: string;

  @Output() added = new EventEmitter<AppModels.CommentModel>();
  profile!: AppModels.MemberModel;
  Content = new FormControl();
  constructor(
    private authService: AuthenticationService,
    private commentService: CommentService,
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
    this.commentService
      .Add({
        Comment: {
          Image: this.image,
          Content: this.Content.value,
          Likes: [],
          PostID: this.postID,
        },
      })
      .subscribe(
        (v) => {
          this.added.emit({
            ID: v.ID,
            PostID: this.postID,
            Content: this.Content.value,
            Image: this.image,
            Likes: [],
          });

          this.Content.setValue('');
          this.image = '';
          this.uploadSuccessful = false;
        },
        (err) => {}
      );
  }
}
