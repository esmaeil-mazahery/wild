export namespace AppModels {
  export class MemberModel {
    public ID?: string;
    public Username!: string;
    public Name!: string;
    public Family!: string;
    public Email!: string;
    public Mobile!: string;
    public Password?: string;
    public Image?: string;
    public Token?: string;
    public ImageHeader?: string;
    public Following?: number = 0;
    public Follower?: number = 0;
    public Biography?: string;
  }

  export class CommentModel {
    public ID?: string;
    public PostID?: string;
    public Content!: string;
    public Image?: string;
    public Likes: string[] = [];

    public Member?: MemberModel;
    public MemberLike?: boolean;
  }

  export class PostModel {
    public ID?: string;
    public Content!: string;
    public Image?: string;
    public Likes: string[] = [];
    public Comments: string[] = [];

    public Member?: MemberModel;
    public MemberLike?: boolean;
  }

  export enum NotifyType {
    NotifyTypeUnknown = 'NotifyTypeUnknown',
    NotifyTypeFollow = 'NotifyTypeFollow',
    NotifyTypeLike = 'NotifyTypeLike',
    NotifyTypeComment = 'NotifyTypeComment',
  }

  export class NotifyModel {
    public ID?: string;
    public Content!: string;
    public Type?: NotifyType;
    public RegisterDate?: string;
    public OwnerMemberID?: string;
    public TargetMemberID?: string;
    public OwnerMember?: MemberModel;
    public TargetMember!: MemberModel;
  }
}
