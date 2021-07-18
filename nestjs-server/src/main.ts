import {NestFactory} from '@nestjs/core';
import {AppModule} from './app.module';
import {NestExpressApplication} from '@nestjs/platform-express';

let path = require('path')

async function bootstrap() {
    const app = await NestFactory.create<NestExpressApplication>(AppModule);
    app.enableCors()
    app.useStaticAssets(path.join(__dirname, '..', 'views'), {
        prefix: '/'
    });
    await app.listen(3000);
}

bootstrap();
