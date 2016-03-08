var gulp = require('gulp');
var gutil = require('gulp-util');
var concat = require('gulp-concat');
var sass = require('gulp-sass');
var minifyCss = require('gulp-minify-css');
var rename = require('gulp-rename');
var coffee = require('gulp-coffee');

var paths = {
  sass: ['public/scss/**/*.scss', 'public/sass/**/*.sass'],
  coffee: ['public/coffee/**/*.coffee']
};

gulp.task('default', ['watch']);

gulp.task('coffee', function() {
  gulp.src('./public/coffee/*.coffee')
    .pipe(coffee({bare: true}).on('error', gutil.log))
    .pipe(gulp.dest('./public/js/'))
});

gulp.task('sass', function(done) {
  gulp.src(paths.sass)
    .pipe(sass({
      indentedSyntax: true,
      errLogToConsole: true
    }))
    .pipe(gulp.dest('./public/css/'))
    .pipe(minifyCss({
      keepSpecialComments: 0
    }))
    .pipe(rename({ extname: '.min.css' }))
    .pipe(gulp.dest('./public/css/'))
    .on('end', done);
});


gulp.task('watch', function() {
  gulp.watch(paths.sass, ['sass']);
  gulp.watch(paths.coffee, ['coffee']);
});
