package model

type CollectionParams struct {
	CollectionUrl   string
	ExternalLink    string
	TwitterLink     string
	InstagramLink   string
	TelegramLink    string
	DiscordLink     string
	LogoImage       string
	FeaturedImage   string
	BannerImage     string
	PaymentAssetIds string
	Description     string
}
type CollectionOption struct {
	F func(*CollectionParams)
}

func WithCollectionUrl(CollectionUrl string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.CollectionUrl = CollectionUrl
	}}
}
func WithExternalLink(ExternalLink string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.ExternalLink = ExternalLink
	}}
}
func WithTwitterLink(TwitterLink string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.TwitterLink = TwitterLink
	}}
}
func WithInstagramLink(InstagramLink string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.InstagramLink = InstagramLink
	}}
}
func WithTelegramLink(TelegramLink string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.TelegramLink = TelegramLink
	}}
}
func WithDiscordLink(DiscordLink string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.DiscordLink = DiscordLink
	}}
}
func WithLogoImage(LogoImage string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.LogoImage = LogoImage
	}}
}
func WithFeaturedImage(FeaturedImage string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.FeaturedImage = FeaturedImage
	}}
}
func WithBannerImage(BannerImage string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.BannerImage = BannerImage
	}}
}
func WithPaymentAssetIds(PaymentAssetIds string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.PaymentAssetIds = PaymentAssetIds
	}}
}
func WithDescription(PaymentAssetIds string) CollectionOption {
	return CollectionOption{func(mp *CollectionParams) {
		mp.PaymentAssetIds = PaymentAssetIds
	}}
}
