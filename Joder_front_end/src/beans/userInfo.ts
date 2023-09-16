export interface User {
  name: string
  userId: string
  picUrl: string
  email: string
  selfIntro: string
  tags: Tag[]
}

export interface MainUser extends User {
  pickedUser: User[]
  matchedUser: User[]
}

export interface Tag {
  content: string
}
