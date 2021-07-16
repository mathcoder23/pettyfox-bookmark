import {Body, Controller, Get, Inject, Post} from '@nestjs/common';
import {Bookmark} from "./bookmark.entity";
import {BookmarkService} from "./bookmark.service";
import {ObjResponse} from "../../comms/api/obj-response";

@Controller('bookmark')
export class BookmarkController {
    constructor(@Inject(BookmarkService) private readonly bookmarkService: BookmarkService) {
    }


    @Post('add')
    add(@Body() bookmark: Bookmark): ObjResponse {
        this.bookmarkService.add(bookmark)
        return ObjResponse.ok(null)
    }

    @Get('list')
    async list(): Promise<ObjResponse> {
        let rep = await this.bookmarkService.list()
        return ObjResponse.ok(rep)
    }
}
