export interface UserInfo{
    userId: string;
    picUrl: string;
    email: string;
    selfIntro: string;
    tags: Tag[];
    pickedUser: OtherUser[];
    matchedUser: OtherUser[];
}

interface OtherUser{
    userId: string;
    picUrl: string;
    email: string;
    selfIntro: string;
    tags: Tag[];
}

export interface Tag{
    catagory: string;
    title: string;
}