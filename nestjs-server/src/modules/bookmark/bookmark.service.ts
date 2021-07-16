import {Injectable} from '@nestjs/common';
import {Bookmark} from "./bookmark.entity";
import {SonicService} from "../sonic/sonic.service";

let md5 = require("md5")
import {RedisService} from "@liaoliaots/nestjs-redis";

@Injectable()
export class BookmarkService {
    private redis: any
    static KEY = "bookmark::user1"

    constructor(private readonly sonicService: SonicService, private redisService: RedisService) {
        this.getRedis()
    }

    private async getRedis() {
        this.redis = await this.redisService.getClient()
    }

    private addBookmarkSonic(bookmark: Bookmark) {
        this.sonicService.addData(BookmarkService.KEY, "url", bookmark.id, bookmark.url)
        if (bookmark.desc && bookmark.desc.length > 0) {
            this.sonicService.addData(BookmarkService.KEY, "desc", bookmark.id, bookmark.desc)
        }
    }

    async search(keyword: string): Promise<Bookmark[]> {
        let a1 = await this.sonicService.searchData(BookmarkService.KEY, "url", keyword)
        console.log('rr', a1)
        let a2 = await this.sonicService.searchData(BookmarkService.KEY, "desc", keyword)
        let results = new Set(a1.concat(a2))
        let rs = []
        console.log('results', results)
        for (let item of results) {
            console.log('item', item)
            rs.push(JSON.parse(await this.redis.hget(BookmarkService.KEY, item)))
        }
        return rs
    }

    add(bookmark: Bookmark): void {
        console.log('b', bookmark)
        if (!bookmark.id) {
            bookmark.id = md5(bookmark.url)
        }
        this.redis.hset(BookmarkService.KEY, bookmark.id, JSON.stringify(bookmark))
        this.addBookmarkSonic(bookmark)
    }

    async list(): Promise<Bookmark[]> {
        let list = await this.redis.hvals(BookmarkService.KEY)
        let temp = []
        for (let item of list) {
            temp.push(JSON.parse(item))
        }
        console.log('list', temp)
        return temp
    }

    remove(ids: string[]): void {
        for (let id of ids) {
            this.redis.hdel(BookmarkService.KEY, id)
        }
    }

    update(bookmark: Bookmark): void {
        if (!bookmark.id) {
            return
        }
        this.redis.hset(BookmarkService.KEY, bookmark.id, JSON.stringify(bookmark))
    }
}
