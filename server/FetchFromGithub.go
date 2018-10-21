package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"sort"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/mmcdole/gofeed"
)

type Follower struct {
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

type Repository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          interface{} `json:"license"`
	Forks            int         `json:"forks"`
	OpenIssues       int         `json:"open_issues"`
	Watchers         int         `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
}

type Article struct {
	RenderedBody   string      `json:"rendered_body"`
	Body           string      `json:"body"`
	Coediting      bool        `json:"coediting"`
	CommentsCount  int         `json:"comments_count"`
	CreatedAt      time.Time   `json:"created_at"`
	Group          interface{} `json:"group"`
	ID             string      `json:"id"`
	LikesCount     int         `json:"likes_count"`
	Private        bool        `json:"private"`
	ReactionsCount int         `json:"reactions_count"`
	Tags           []struct {
		Name     string        `json:"name"`
		Versions []interface{} `json:"versions"`
	} `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	User      struct {
		Description       interface{} `json:"description"`
		FacebookID        interface{} `json:"facebook_id"`
		FolloweesCount    int         `json:"followees_count"`
		FollowersCount    int         `json:"followers_count"`
		GithubLoginName   interface{} `json:"github_login_name"`
		ID                string      `json:"id"`
		ItemsCount        int         `json:"items_count"`
		LinkedinID        interface{} `json:"linkedin_id"`
		Location          interface{} `json:"location"`
		Name              string      `json:"name"`
		Organization      interface{} `json:"organization"`
		PermanentID       int         `json:"permanent_id"`
		ProfileImageURL   string      `json:"profile_image_url"`
		TwitterScreenName string      `json:"twitter_screen_name"`
		WebsiteURL        interface{} `json:"website_url"`
	} `json:"user"`
	PageViewsCount interface{} `json:"page_views_count"`
}

type hatenaArticle struct {
	Title string `json:"title"`
	Link  string `json:"link""`
}

type hatenaArticles []hatenaArticle

type Repositories []Repository

type Articles []Article

func (r Repositories) Len() int {
	return len(r)
}

func (r Repositories) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Repositories) Less(i, j int) bool {
	return r[i].StargazersCount < r[j].StargazersCount
}

func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Articles) Less(i, j int) bool {
	return a[i].LikesCount < a[j].LikesCount
}

func fetchUserFromGithub(username string) ([]Follower, error) {
	url := "https://api.github.com/users/" + username + "/followers"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	var articles []Follower
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func fetchRepositoriesFromGithub(username string) (Repositories, error) {
	url := "https://api.github.com/users/" + username + "/repos"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	var response Repositories
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func fetchArticlesFromQiita(username string) (Articles, error) {
	url := "https://qiita.com/api/v2/users/" + username + "/items"
	//url := "https://qiita.com/api/v2/users/yuno_miyako/items"
	fmt.Println(username)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	var response Articles
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		c.Next()
	}
}

func KusaHandler(c *gin.Context) {
	name := c.Param("name")
	res, err := http.Get("https://github.com/users/" + name + "/contributions")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items

	svg := doc.Find("svg").SetAttr("xmlns", "http://www.w3.org/2000/svg").AppendHtml(`<style>text{font-size: 9px;fill: #767676;}</style>`)
	html, err := svg.Parent().Html()
	if err != nil {
		log.Fatal(err)
	}
	c.Writer.Header().Set("Content-Type", "image/svg+xml")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.String(200, html)
	return
}

func RepositoriesHandler(c *gin.Context) {
	name := c.Param("name")
	repositories, err := fetchRepositoriesFromGithub(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	sort.Sort(sort.Reverse(repositories))

	var buf bytes.Buffer
	b, _ := json.Marshal(repositories)
	buf.Write(b)
	c.String(http.StatusOK, buf.String())
}
func ArticlesHandler(c *gin.Context) {
	name := c.Param("name")
	articles, err := fetchArticlesFromQiita(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	sort.Sort(sort.Reverse(articles))

	var buf bytes.Buffer
	b, _ := json.Marshal(articles)
	buf.Write(b)
	c.String(http.StatusOK, buf.String())
}
func hatenaArticleHandler(c *gin.Context) {
	want_url := c.Param("url")

	fp := gofeed.NewParser()
	url := "https://" + want_url + "/feed"
	feed, _ := fp.ParseURL(url)
	items := feed.Items
	var articles hatenaArticles

	for _, item := range items {
		var article hatenaArticle = hatenaArticle{
			Title: item.Title,
			Link:  item.Link,
		}
		articles = append(articles, article)

	}
	fmt.Println(articles)
	var buf bytes.Buffer
	b, _ := json.Marshal(articles)
	buf.Write(b)
	c.String(http.StatusOK, buf.String())
}

func main() {
	router := gin.Default()
	//CORS対策
	router.Use(CORSMiddleware())

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		follower, err := fetchUserFromGithub(name)
		if err != nil {
			//落ちると困るのでJSONでエラーを吐く
			//log.Fatalf("Error!: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}

		//構造体をJSONに戻す
		var buf bytes.Buffer
		b, _ := json.Marshal(follower)
		buf.Write(b)
		c.String(http.StatusOK, buf.String())
	})

	router.GET("/user/:name/kusa", KusaHandler)
	router.GET("/repos/user/:name", RepositoriesHandler)
	router.GET("/articles/user/:name", ArticlesHandler)
	router.GET("/hatenaarticles/url/:url", hatenaArticleHandler)
	router.Run(":8080")
}
