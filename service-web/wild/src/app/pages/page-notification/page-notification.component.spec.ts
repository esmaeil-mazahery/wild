import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PageNotificationComponent } from './page-notification.component';

describe('PageNotificationComponent', () => {
  let component: PageNotificationComponent;
  let fixture: ComponentFixture<PageNotificationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PageNotificationComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PageNotificationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
