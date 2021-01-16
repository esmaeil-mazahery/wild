import { PostService } from 'projects/libs/infrastructure/src/serivces/Entity/post.service';
import { AlertService } from './../../../../projects/libs/infrastructure/src/serivces/system/alert.service';
import { UploadService } from './../../../../projects/libs/infrastructure/src/serivces/Entity/upload.service';
import {
  AuthenticationModels,
  AuthenticationService,
} from 'projects/libs/infrastructure/src/serivces/Entity/auth.service';
import { BaseComponent } from 'projects/libs/infrastructure/src/components/base-component/base.component';
import { Component, OnInit, ViewChild } from '@angular/core';
import { forkJoin, Observable } from 'rxjs';
import { AppModels } from 'projects/libs/infrastructure/src/serivces/Entity/models';

@Component({
  templateUrl: './page-profile.component.html',
  styleUrls: ['./page-profile.component.scss'],
})
export class PageProfileComponent extends BaseComponent implements OnInit {
  profile!: AppModels.MemberModel;
  postList: AppModels.PostModel[] = [];
  constructor(
    private authService: AuthenticationService,
    private postService: PostService,
    private uploadService: UploadService,
    private alertService: AlertService
  ) {
    super();
  }

  ngOnInit(): void {
    this.authService.GetProfileSync().subscribe((v) => {
      this.profile = v;
    });

    this.loadPosts();
  }

  currentPage: number = 1;
  loadPosts() {
    this.postService
      .MyPosts({
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

  @ViewChild('file') file: any;
  uploading = false;
  uploadSuccessful = false;
  public files: Set<File> = new Set();
  address: string | undefined;
  progress:
    | {
        [key: string]: { progress: Observable<number>; filename: string };
      }
    | undefined;
  progressNumber: number = 0;

  isImageProfile: boolean = true;

  addFiles(isImageProfile: boolean) {
    this.isImageProfile = isImageProfile;
    this.file.nativeElement.click();
  }

  onFilesAdded() {
    const files: { [key: string]: File } = this.file.nativeElement.files;
    for (const key in files) {
      if (!isNaN(parseInt(key, 10))) {
        this.uploadSuccessful = false;
        this.address = files[key].name;
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
          if (this.isImageProfile) {
            this.authService
              .ChangeImageProfile({
                ImageURL: this.progress[key].filename,
              })
              .subscribe(
                (v) => {
                  this.alertService.openSnackBar(
                    'تصویر با موفقیت تغییر یافت',
                    false,
                    3000
                  );
                },
                (err) => {
                  this.alertService.openSnackBar('خطایی رخ داد', false, 3000);
                }
              );
          } else {
            this.authService
              .ChangeImageHeader({
                ImageURL: this.progress[key].filename,
              })
              .subscribe(
                (v) => {
                  this.alertService.openSnackBar(
                    'تصویر با موفقیت تغییر یافت',
                    false,
                    3000
                  );
                },
                (err) => {
                  this.alertService.openSnackBar('خطایی رخ داد', false, 3000);
                }
              );
          }
        }
      }

      // ... the upload was successful...
      this.uploadSuccessful = true;

      // ... and the component is no longer uploading
      this.uploading = false;
      this.progressNumber = 0;
    });
  }

  removeImage(isImageProfile: boolean) {
    if (this.isImageProfile) {
      this.authService
        .ChangeImageProfile({
          ImageURL: '',
        })
        .subscribe(
          (v) => {
            this.alertService.openSnackBar(
              'تصویر با موفقیت حذف شد',
              false,
              3000
            );
          },
          (err) => {
            this.alertService.openSnackBar('خطایی رخ داد', false, 3000);
          }
        );
    } else {
      this.authService
        .ChangeImageHeader({
          ImageURL: '',
        })
        .subscribe(
          (v) => {
            this.alertService.openSnackBar(
              'تصویر با موفقیت حذف شد',
              false,
              3000
            );
          },
          (err) => {
            this.alertService.openSnackBar('خطایی رخ داد', false, 3000);
          }
        );
    }
  }
}
