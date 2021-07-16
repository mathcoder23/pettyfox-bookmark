import {Module} from '@nestjs/common';
import {SonicService} from './sonic.service';

@Module({
    providers: [SonicService],
    exports: [SonicService]
})
export class SonicModule {
}
