package vk

type Updates struct {
	Updates []struct {
		Type   string
		Object struct {
			MessageNew Message `json:"message"`
			Message
			Wall        Wall
			WallComment WallComment
		}
	}
}

type Message struct {
	Text         string
	PeerId       int `json:"peer_id"`
	Id           int
	FromId       int `json:"from_id"`
	Date         int
	RandomId     int `json:"random_id"`
	Important    bool
	Payload      string
	ReplyMessage *Message `json:"reply_message"`
	Keyboard     Keyboard
	Ref          string
	RefSource    string `json:"ref_source"`
	Action       struct {
		Type     string
		MemberId int `json:"member_id"`
		Text     string
		Email    string
		Photo    struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		}
	}
	Geo struct {
		Type        string
		Coordinates struct {
			Latitude, Longitude float32
		}
		Place Place
	}
	ConversationMessageId int `json:"conversation_message_id"`
	Attachments           []Attachments
	FwdMessages           []*Message `json:"fwd_messages"`
}

type User struct {
	// Main Fields
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Deactivated     string `json:"deactivated"`
	IsClosed        bool   `json:"is_closed"`
	CanAccessClosed bool   `json:"can_access_closed"`
	// Extra Fields
	About                  string `json:"about"`
	Activities             string `json:"activities"`
	BDate                  string `json:"bdate"`
	Blacklisted            int    `json:"blacklisted"`
	BlacklistedByMe        int    `json:"blacklisted_by_me"`
	Books                  string `json:"books"`
	CanPost                int    `json:"can_post"`
	CanSeeAllPosts         int    `json:"can_see_all_posts"`
	CanSeeAudio            int    `json:"can_see_audio"`
	CanSendFriendRequest   int    `json:"can_send_friend_request"`
	CanWritePrivateMessage int    `json:"can_write_private_message"`
	Career                 struct {
		GroupID   int    `json:"group_id"`
		Company   string `json:"company"`
		CountryID int    `json:"country_id"`
		CityID    int    `json:"city_id"`
		CityName  string `json:"city_name"`
		From      int    `json:"from"`
		Until     int    `json:"until"`
		Position  string `json:"position"`
	} `json:"career"`
	City struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	CommonCount int `json:"common_count"`
	Connections struct {
		Instagram   string `json:"instagram"`
		Skype       string `json:"skype"`
		Facebook    string `json:"facebook"`
		Twitter     string `json:"twitter"`
		LiveJournal string `json:"livejournal"`
	} `json:"connections"`
	Contacts struct {
		MobilePhone string `json:"mobile_phone"`
		HomePhone   string `json:"home_phone"`
	} `json:"contacts"`
	Counters struct {
		Albums, Videos, Audio, Photos, Notes, Friends, Groups, Followers, Pages int
		OnlineFriends                                                           int `json:"online_friends"`
		MutualFriends                                                           int `json:"mutual_friends"`
		UserVideos                                                              int `json:"user_videos"`
	} `json:"counters"`
	Country struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
	CropPhoto struct {
		Photo Photo `json:"photo"`
	} `json:"crop_photo"`
	Domain    string `json:"domain"`
	Education struct {
		University     int    `json:"university"`
		Faculty        int    `json:"faculty"`
		Graduation     int    `json:" "`
		UniversityName string `json:"university_name"`
		FacultyName    string `json:"faculty_name"`
	} `json:"education"`
	FollowersCount   int    `json:"followers_count"`
	FriendStatus     int    `json:"friend_status"`
	Games            string `json:"games"`
	HasMobile        int    `json:"has_mobile"`
	HasPhoto         int    `json:"has_photo"`
	HomeTown         string `json:"home_town"`
	Interests        string `json:"interests"`
	IsFavorite       int    `json:"is_favorite"`
	IsFriend         int    `json:"is_friend"`
	IsHiddenFromFeed int    `json:"is_hidden_from_feed"`
	LastNameNom      string `json:"last_name_nom"`
	LastNameGen      string `json:"last_name_gen"`
	LastNameDat      string `json:"last_name_dat"`
	LastNameAcc      string `json:"last_name_acc"`
	LastNameIns      string `json:"last_name_ins"`
	LastNameAbl      string `json:"last_name_abl"`
	LastSeen         struct {
		Time     int `json:"time"`
		Platform int `json:"platform"`
	} `json:"last_seen"`
	Lists      string `json:"lists"`
	MaidenName string `json:"maiden_name"`
	Military   struct {
		Unit      string
		UnitId    int `json:"unit_id"`
		CountryId int `json:"country_id"`
		From      int `json:"from"`
		Until     int `json:"until"`
	} `json:"military"`
	Movies     string `json:"movies"`
	Music      string `json:"music"`
	Nickname   string `json:"nickname"`
	Occupation struct {
		Type string
		ID   int
		Name string
	} `json:"occupation"`
	Online       int `json:"online"`
	OnlineMobile int `json:"online_mobile"`
	OnlineApp    int `json:"online_app"`
	Personal     struct {
		Political  int    `json:"political"`
		Langs      []int  `json:"langs"`
		Religion   string `json:"religion"`
		InspiredBy string `json:"inspired_by"`
		PeopleMain int    `json:"people_main"`
		LifeMain   int    `json:"life_main"`
		Smoking    int    `json:"smoking"`
		Alcohol    int    `json:"alcohol"`
	} `json:"personal"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200Orig string `json:"photo_200_orig"`
	Photo200     string `json:"photo_200"`
	Photo400Orig string `json:"photo_400_orig"`
	PhotoId      string `json:"photo_id"`
	PhotoMax     string `json:"photo_max"`
	PhotoMaxOrig string `json:"photo_max_orig"`
	Quotes       string `json:"quotes"`
	Relatives    []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"` //Child Sibling Parent GrandParent GrandChild
	} `json:"relatives"`
	Relation int `json:"relation"`
	Schools  []struct {
		Id, Country, City int
		Name              string `json:"name"`
		YearFrom          int    `json:"year_from"`
		YearTo            int    `json:"year_to"`
		YearGraduated     int    `json:"year_graduated"`
		Class             string
		Speciality        string
		Type              int    // Тип школы
		TypeStr           string `json:"type_str"` //0 — "школа"; 1 — "гимназия"; 2 —"лицей"; 3 — "школа-интернат"; 4 — "школа вечерняя"; 5 — "школа музыкальная"; 6 — "школа спортивная"; 7 — "школа художественная"; 8 — "колледж"; 9 — "профессиональный лицей"; 10 — "техникум"; 11 — "ПТУ"; 12 — "училище"; 13 — "школа искусств".
	}
	ScreenName   string `json:"screen_name"`
	Sex          int
	Site         string
	Timezone     int
	Trending     int
	Tv           string
	Universities []struct {
		Id, Country, City int
		Name              string
		Faculty           int
		FacultyName       string `json:"faculty_name"`
		Chair             int
		ChairName         string `json:"chair_name"`
		Graduation        int
		EducationForm     string `json:"education_form"`
		EducationStatus   string `json:"education_status"`
	}
	Verified    int
	WallDefault string `json:"wall_default"`
}
type Group struct {
	// Main
	Id           int
	Name         string
	ScreenName   string `json:"screen_name"`
	IsClosed     int    `json:"is_closed"`
	Deactivated  string
	IsAdmin      int `json:"is_admin"`
	AdminLevel   int `json:"admin_level"`
	IsMember     int `json:"is_member"`
	IsAdvertiser int `json:"is_advertiser"`
	InvitedBy    int `json:"invited_by"`
	Type         string
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`
	// Extra
	Activity          string
	AgeLimits         int `json:"age_limits"`
	CanCreateTopic    int `json:"can_create_topic"`
	CanMessage        int `json:"can_message"`
	CanPost           int `json:"can_post"`
	CanSeeAllPosts    int `json:"can_see_all_posts"`
	CanUploadDoc      int `json:"can_upload_doc"`
	CanUploadVideo    int `json:"can_upload_video"`
	Description       string
	FixedPost         int `json:"fixed_post"`
	HasPhoto          int `json:"has_photo"`
	IsFavorite        int `json:"is_favorite"`
	IsHiddenFromFeed  int `json:"is_hidden_from_feed"`
	IsMessagesBlocked int `json:"is_messages_blocked"`
	MainAlbumId       int `json:"main_album_id"`
	MainSection       int `json:"main_section"`
	MemberStatus      int `json:"member_status"`
	MembersCount      int `json:"members_count"`
	Place             Place
	PublicDateLabel   string `json:"public_date_label"`
	Site              string
	StartDate         int `json:"start_date"`
	FinishDate        int `json:"finish_date"`
	Status            string
	Trending          int
	Verified          int
	Wall              int
	WikiPage          string `json:"wiki_page"`
	Addresses         struct {
		IsEnabled     bool `json:"is_enabled"`
		MainAddressId int  `json:"main_address_id"`
	}
	BanInfo struct {
		EndDate int `json:"end_date"`
		Comment string
	} `json:"ban_info"`
	City struct {
		Id    int
		Title string
	}
	Contacts []struct {
		UserId int `json:"user_id"`
		Desc   string
		Phone  string
		Email  string
	}
	Counters struct {
		Photos,
		Albums,
		Audios,
		Videos,
		Topics,
		Docs int
	}
	Country struct {
		Id    int
		Title string
	}
	Cover struct {
		Enabled int
		Images  []struct {
			Url    string
			Width  int
			Height int
		}
	}
	CropPhoto struct {
		Photo Photo
		Crop  struct {
			X, Y, X1, Y1 float32
		} `json:"crop_photo"`
		Rect struct {
			X, Y, X1, Y1 float32
		}
	}
	Links []struct {
		Id       int
		Url      string
		Name     string
		Desc     string
		Photo50  string `json:"photo_50"`
		Photo100 string `json:"photo_100"`
	}
	Market struct {
		Enabled      int
		PriceMin     int `json:"price_min"`
		PriceMax     int `json:"price_max"`
		MainAlbumId  int `json:"main_album_id"`
		ContactId    int `json:"contact_id"`
		Currency     Currency
		CurrencyText string `json:"currency_text"`
	}
}
type Conversation struct {
	InRead      int `json:"in_read"`
	OutRead     int `json:"out_read"`
	UnreadCount int `json:"unread_count"`
	Important   bool
	Unanswered  bool
	Peer        struct {
		Id      int
		Type    string
		LocalId int `json:"local_id"`
	}
	PushSettings struct {
		DisabledUntill  int  `json:"disabled_untill"`
		DisabledForever bool `json:"disabled_forever"`
		NoSound         bool `json:"no_sound"`
	} `json:"push_settings"`
	CanWrite struct {
		Allowed bool
		Reason  int
	} `json:"can_write"`
	ChatSettings struct {
		MembersCount  int `json:"members_count"`
		Title         string
		PinnedMessage Message `json:"pinned_message"`
		State         string
		Photo         struct {
			Photo50  string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		}
		ActiveIds      []int `json:"active_ids"`
		IsGroupChannel bool  `json:"is_group_channel"`
	} `json:"chat_settings"`
}
type Chat struct {
	Id           int
	Type         string
	AdminId      int `json:"admin_id"`
	Users        []int
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`
	Left         int
	Kicked       int
	PushSettings struct {
		Sound         int
		DisabledUntil int `json:"disabled_until"`
	} `json:"push_settings"`
}

type Attachments struct {
	Type         string
	Photo        Photo
	Audio        Audio
	Video        Video
	Sticker      Sticker
	Wall         Wall
	Note         Note
	Market       Market
	MarketAlbum  MarketAlbum `json:"market_album"`
	Page         Page
	App          App
	Pool         Pool
	WallComment  WallComment
	Graffiti     Graffiti
	AudioMessage AudioMessage `json:"audio_message"`
	Document     Document
}
type Photo struct {
	Id            int
	OwnerId       int `json:"owner_id"`
	Date          int
	UserId        int `json:"user_id"`
	AlbumId       int `json:"album_id"`
	Sizes         []PhotoSizes
	Width, Height int
}
type PhotoSizes struct {
	Type, Url     string
	Width, Height int
}
type Audio struct {
	Artist   string
	Title    string
	AlbumId  int `json:"album_id"`
	Id       int
	OwnerId  int `json:"owner_id"`
	Duration int
	Url      string
	LyricsId int `json:"lyrics_id"`
	GenreId  int `json:"genre_id"`
	Date     int
	NoSearch int `json:"no_search"`
	IsHq     int `json:"is_hq"`
}
type Video struct {
	Id             int
	OwnerId        int `json:"owner_id"`
	Title          string
	Description    string
	Duration       int
	Photo130       string `json:"photo_130"`
	Photo320       string `json:"photo_320"`
	Photo640       string `json:"photo_640"`
	Photo800       string `json:"photo_800"`
	Photo1280      string `json:"photo_1280"`
	FirstFrame130  string `json:"first_frame_130"`
	FirstFrame320  string `json:"first_frame_320"`
	FirstFrame640  string `json:"first_frame_640"`
	FirstFrame800  string `json:"first_frame_800"`
	FirstFrame1280 string `json:"first_frame_1280"`
	Date           int
	AddingDate     int `json:"adding_date"`
	Views          int
	Comments       int
	Player         string
	Platform       string
	CanEdit        int    `json:"can_edit"`
	CanAdd         int    `json:"can_add"`
	IsPrivate      int    `json:"is_private"`
	AccessKey      string `json:"access_key"`
	Processing     int
	Live           int
	Upcoming       int
	IsFavorite     bool `json:"is_favorite"`
}
type Document struct {
	Id      int
	OwnerId int `json:"owner_id"`
	Title   string
	Size    int
	Ext     string
	Url     string
	Date    int
	Type    int
	Preview struct {
		Graffiti     Graffiti
		AudioMessage AudioMessage `json:"audio_message"`
		Photo        struct {
			Sizes PhotoSizes
		}
	}
}
type Graffiti struct {
	Src    string
	Width  int
	Height int
}
type AudioMessage struct {
	Duration int
	Waveform []int
	LinkOgg  string `json:"link_ogg"`
	LinkMp3  string `json:"link_mp_3"`
}
type Link struct {
	Url         string
	Title       string
	Caption     string
	Description string
	Photo       Photo
	Product     Product
}
type Product struct {
	Price struct {
		Amount   int
		Currency Currency
		Text     string
	}
}
type Currency struct {
	Id   int
	Name string
}
type Sticker struct {
	ProductId int `json:"product_id"`
	StickerId int `json:"sticker_id"`
	Images    []struct {
		Url           string
		Width, Height int
	}
	ImagesWithBackground []struct {
		Url           string
		Width, Height int
	} `json:"images_with_background"`
}
type Wall struct {
	Id           int
	OwnerId      int `json:"owner_id"`
	FromId       int `json:"from_id"`
	CreatedBy    int `json:"created_by"`
	Date         int
	Text         string
	ReplyOwnerId int `json:"reply_owner_id"`
	ReplyPostId  int `json:"reply_post_id"`
	FriendsOnly  int `json:"friends_only"`
	Comments     struct {
		Count         int
		CanPost       int  `json:"can_post"`
		GroupsCanPost int  `json:"groups_can_post"`
		CanClose      bool `json:"can_close"`
		CanOpen       bool `json:"can_open"`
	}
	Copyright string
	Likes     struct {
		Count      int
		UserLikes  int `json:"user_likes"`
		CanLike    int `json:"can_like"`
		CanPublish int `json:"can_publish"`
	}
	Reposts struct {
		Count        int
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Views struct {
		Count int
	}
	PostType   string `json:"post_type"`
	PostSource struct {
		Type     string
		Platform string
		Data     string
		Url      string
	} `json:"post_source"`
	Attachments []Attachments
	Geo         struct {
		Type        string
		Coordinates string
		Place       Place
	}
	SignerId    int  `json:"signer_id"`
	CanPin      int  `json:"can_pin"`
	CanDelete   int  `json:"can_delete"`
	CanEdit     int  `json:"can_edit"`
	IsPinned    int  `json:"is_pinned"`
	MarkedAsAds int  `json:"marked_as_ads"`
	IsFavorite  bool `json:"is_favorite"`
	PostponedId int  `json:"postponed_id"`
}
type WallComment struct {
	Id             int
	FromId         int `json:"from_id"`
	Date           int
	Text           string
	ReplyToUser    int `json:"reply_to_user"`
	ReplyToComment int `json:"reply_to_comment"`
	Attachments    []Attachments
	ParentsStack   []int `json:"parents_stack"`
	Thread         struct {
		Count           int
		Items           []*WallComment
		CanPost         bool `json:"can_post"`
		ShowReplyButton bool `json:"show_reply_button"`
		GroupsCanPost   bool `json:"groups_can_post"`
	}
}
type Address struct {
	Id                int
	CountryId         int `json:"country_id"`
	CityId            int `json:"city_id"`
	Title             string
	Address           string
	AdditionalAddress string `json:"additional_address"`
	Latitude          float32
	Longitude         float32
	MetroStationId    int    `json:"metro_station_id"`
	WorkInfoStatus    string `json:"work_info_status"`
	Timetable         struct {
		Fri GroupTimetable
		Mon GroupTimetable
		Sat GroupTimetable
		Thu GroupTimetable
		Tue GroupTimetable
		Wed GroupTimetable
	}
}
type GroupTimetable struct {
	CloseTime      int `json:"close_time"`
	OpenTime       int `json:"open_time"`
	BreakCloseTime int `json:"break_close_time"`
	BreakOpenTime  int `json:"break_open_time"`
}
type Note struct {
	Id           int
	OwnerId      int `json:"owner_id"`
	Title        string
	Text         string
	Date         int
	Comments     int
	ReadComments int `json:"read_comments"`
	ViewUrl      int `json:"view_url"`
}
type Page struct {
	Id                       int
	GroupId                  int `json:"group_id"`
	CreatorId                int `json:"creator_id"`
	Title                    string
	CurrentUserCanEdit       int `json:"current_user_can_edit"`
	CurrentUserCanEditAccess int `json:"current_user_can_edit_access"`
	WhoCanView               int `json:"who_can_view"`
	WhoCanEdit               int `json:"who_can_edit"`
	Edited                   int
	Created                  int
	EditorId                 int `json:"editor_id"`
	Views                    int
	Parent                   string
	Parent2                  string
	Source                   string
	Html                     string
	ViewUrl                  string `json:"view_url"`
}
type Market struct {
	// Main
	Id           int
	OwnerId      int `json:"owner_id"`
	Title        string
	Description  string
	ThumbPhoto   string `json:"thumb_photo"`
	Date         int
	Availability int
	IsFavorite   bool `json:"is_favorite"`
	Price        struct {
		Amount    int
		OldAmount int `json:"old_amount"`
		Text      string
		Currency  Currency
	}
	Category struct {
		Id      int
		Name    string
		Section struct {
			Id   int
			Name string
		}
	}
	// Extra
	Photos      []Photo
	CanComment  int `json:"can_comment"`
	CanRepost   int `json:"can_repost"`
	Url         string
	ButtonTitle string `json:"button_title"`
	Likes       struct {
		UserLikes int `json:"user_likes"`
		Count     int
	}
}
type MarketAlbum struct {
	Id          int
	OwnerId     int `json:"owner_id"`
	Title       string
	Photo       Photo
	Count       int
	UpdatedTime int `json:"updated_time"`
}
type Order struct {
	Id                 int
	GroupId            int `json:"group_id"`
	UserId             int `json:"user_id"`
	Date               int
	VariantsGroupingId int  `json:"variants_grouping_id"`
	IsMainVariant      bool `json:"is_main_variant"`
	CartQuantity       int  `json:"cart_quantity"`
	Status             int
	ItemsCount         int    `json:"items_count"`
	DisplayOrderId     string `json:"display_order_id"`
	TrackNumber        string `json:"track_number"`
	TrackLink          string `json:"track_link"`
	Comment            string
	Address            string
	PreviewOrderItems  []Market `json:"preview_order_items"`
	PropertyValues     []struct {
		VariantId    int    `json:"variant_id"`
		VariantName  string `json:"variant_name"`
		PropertyName string `json:"property_name"`
	} `json:"property_values"`
	TotalPrice struct {
		Amount   string
		Text     string
		Currency Currency
	} `json:"total_price"`
}
type Board struct {
	Id           int
	Title        string
	Created      int
	CreatedBy    int `json:"created_by"`
	Updated      int
	IsClosed     int `json:"is_closed"`
	IsFixed      int `json:"is_fixed"`
	Comments     int
	FirstComment string `json:"first_comment"`
	LastComment  string `json:"last_comment"`
}
type BoardComment struct {
	Id          int
	FromId      int `json:"from_id"`
	Date        int
	Text        string
	Attachments []Attachments
	Likes       struct {
		Count     int
		UserLikes int `json:"user_likes"`
		CanLike   int `json:"can_like"`
	}
}
type App struct {
	// Main
	Id                    int
	Title                 string
	Icon278               string `json:"icon_278"`
	Icon139               string `json:"icon_139"`
	Icon150               string `json:"icon_150"`
	Icon75                string `json:"icon_75"`
	Banner560             string `json:"banner_560"`
	Banner1120            string `json:"banner_1120"`
	Type                  string
	Section               string
	AuthorUrl             string `json:"author_url"`
	AuthorId              string `json:"author_id"`
	AuthorGroup           string `json:"author_group"`
	MembersCount          string `json:"members_count"`
	PublishedDate         int    `json:"published_date"`
	CatalogPosition       int    `json:"catalog_position"`
	International         int
	LeaderboardType       int `json:"leaderboard_type"`
	GenreId               int `json:"genre_id"`
	Genre                 string
	PlatformId            string `json:"platform_id"`
	IsInCatalog           int    `json:"is_in_catalog"`
	Friends               []int
	Installed             int
	IsHtml5App            int `json:"is_html_5_app"`
	ScreenOrientation     int `json:"screen_orientation"`
	MobileControlsType    int `json:"mobile_controls_type"`
	MobileViewSupportType int `json:"mobile_view_support_type"`
	// Extra
	Description string
	ScreenName  string `json:"screen_name"`
	Icon16      string `json:"icon_16"`
	Screenshots []Photo
	PushEnabled int `json:"push_enabled"`
}
type Pool struct {
	Id         int
	OwnerId    int `json:"owner_id"`
	Created    int
	Question   string
	Votes      int
	Anonymous  bool
	Multiple   bool
	AnswerIds  []int `json:"answer_ids"`
	EndDate    int   `json:"end_date"`
	Closed     bool
	IsBoard    bool `json:"is_board"`
	CanEdit    bool `json:"can_edit"`
	CanReport  bool `json:"can_report"`
	CanShare   bool `json:"can_share"`
	AuthorId   int  `json:"author_id"`
	Photo      Photo
	Friends    []int
	Background struct {
		Id     int
		Type   string
		Angle  int
		Color  string
		Width  int
		Height int
		Images []Photo
		Points []struct {
			Position float32
			Color    string
		}
	}
	Answers []struct {
		Id    int
		Text  string
		Votes int
		Rate  float32
	}
}
type Stats struct {
	PeriodFrom int `json:"period_from"`
	PeriodTo   int `json:"period_to"`
	Visitors   struct {
		Views    int
		Visitors int
	}
	Reach struct {
		Reach            int
		ReachSubscribers int `json:"reach_subscribers"`
		MobileReach      int `json:"mobile_reach"`
		Sex              []struct {
			Value string
			Count int
		}
		Age []struct {
			Value string
			Count int
		}
		SexAge []struct {
			Value string
			Count int
		} `json:"sex_age"`
		Cities []struct {
			Name      string
			CityId    interface{} `json:"city_id"`
			CityIdInt int
			Count     int
		}
		Countries []struct {
			Name      string
			Code      string
			CountryId int `json:"country_id"`
			Count     int
		}
	}
}

type Place struct {
	Id                  int
	Title               string
	Latitude, Longitude float32
	Created             int
	Icon                string
	Country             string
	City                string
}
type Keyboard struct {
	OneTime bool `json:"one_time"`
	Inline  bool
	Buttons [][]Button
}
type Button struct {
	Action struct {
		Type    string
		Label   string
		Payload string
		Link    string
		Hash    string
		AppId   int `json:"app_id"`
		OwnerId int `json:"owner_id"`
	}
	Color string
}

type Params map[string]interface{}

type vkError struct {
	Error struct {
		ErrorCode     int    `json:"error_code"`
		ErrorMsg      string `json:"error_msg"`
		RequestParams []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"request_params"`
	} `json:"error"`
}
