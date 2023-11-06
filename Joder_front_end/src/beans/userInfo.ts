export interface User {
  name?: string
  userId?: string
  picUrl?: string[]
  email?: string
  gender?: number
  lastLoginTime?: string
  userType?: string
  tags?: Tag[]
}

export interface InsertUser {
  name?: string
  email?: string
  selfIntro?: string
  userType?: number
  birthday?: string
  gender?: number
}

export interface MainUser extends User {
  unmatchedPaginateNum?: number
  pickedUser?: User[]
  matchedUser?: User[]
}

export interface Tag {
  content: string
}
