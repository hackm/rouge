var gulp = require('gulp');
var riot = require('gulp-riot');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');

gulp.task('riot', function() {
    return gulp.src('src/**/*.tag.html')
        .pipe(riot())
        .pipe(gulp.dest('build'));
});
gulp.task('compress',['riot'], function () {
    return gulp.src('build/**/*.js')
        .pipe(uglify())
        .pipe(concat('tags.js'))
        .pipe(gulp.dest('../../statics/themes/simple'));
});
gulp.task('watch', function () {
    gulp.watch('src/**/*.tag.html', ['compress']);
})
gulp.task('default', ['compress']);