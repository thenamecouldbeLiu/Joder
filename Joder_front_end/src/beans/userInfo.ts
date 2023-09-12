export interface UserInfo{
    userId: string;
    picUrl: string;
    email: string;
    selfIntro: string;
    tags: Tag[];
    pickedUser: UserInfo[];
    matchedUser: UserInfo[];
}

export interface Tag{
    catagory: string;
    title: string;
}