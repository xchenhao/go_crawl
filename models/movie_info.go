package models

import (
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"strings"
)

const (
	KEY_DOUBAN_MOVIE_URL_QUEUE = "movie_url_queue"
	KEY_DOUBAN_MOVIE_URL_VISIT_SET = "movie_url_visited_set"
)

type MovieInfo struct {
  Id int64
  Movie_id int64
  Movie_name string
  Movie_pic string
  Movie_director string
  Movie_writer string
  Movie_country string
  Movie_language string
  Movie_main_character string
  Movie_type string
  Movie_on_time string
  Movie_span string
  Movie_grade string
}

func AddMovie(sMovieHtml string) (int64, error) {
	movieName := GetMovieName(sMovieHtml)
	if "" == movieName {
		return 0, nil
	}
	// 记录电影信息
	var model MovieInfo
	model.Movie_name            = movieName
	model.Movie_director        = GetMovieDirector(sMovieHtml)
	model.Movie_main_character  = GetMovieMainCharacters(sMovieHtml)
	model.Movie_type            = GetMovieGenre(sMovieHtml)
	model.Movie_on_time         = GetMovieOnTime(sMovieHtml)
	model.Movie_grade           = GetMovieGrade(sMovieHtml)
	model.Movie_span            = GetMovieRunningTime(sMovieHtml)
	model.Id = 0

	id, err := GetConnection(new(MovieInfo), "default2").Insert(&model)
	return id, err
}

func GetMovieDirector(movieHtml string) string{
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieName(movieHtml string)string{
	if movieHtml == ""{
		return ""
	}

	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieMainCharacters(movieHtml string)string{
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	mainCharacters := ""
	for _,v := range result{
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")
}

func GetMovieGrade(movieHtml string)string{
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}
	return string(result[0][1])
}

func GetMovieGenre(movieHtml string)string{
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	movieGenre := ""
	for _,v := range result{
		movieGenre += v[1] + "/"
	}
	return strings.Trim(movieGenre, "/")
}

func GetMovieOnTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieRunningTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}


func GetMovieUrls(movieHtml string)[]string{
	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	var movieSets []string
	for _,v := range result{
		movieSets = append(movieSets, v[1])
	}

	return movieSets
}
