package tgin

import (
	"github.com/gotd/td/bin"
	"github.com/gotd/td/tg"
)

type ChatFolderInfo struct {
	Title        string              `json:"title"`
	Emoticon     string              `json:"emoticon"`
	FolderId     int                 `json:"folderId"`
	Delete       bool                `json:"delete"`
	Groups       bool                `json:"groups"`
	NonContacts  bool                `json:"nonContacts"`
	Contacts     bool                `json:"contacts"`
	Channels     bool                `json:"channels"`
	Bots         bool                `json:"bots"`
	IncludePeers []tg.InputPeerClass `json:"includePeers"`
	PinnedPeers  []tg.InputPeerClass `json:"pinnedPeers"`
	ExcludePeers []tg.InputPeerClass `json:"excludePeers"`
}

type DialogFilter struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `json:"flags,omitempty"`
	// Whether to include all contacts in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Contacts bool `json:"contacts,omitempty"`
	// Whether to include all non-contacts in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	NonContacts bool `json:"nonContacts,omitempty"`
	// Whether to include all groups in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Groups bool `json:"groups,omitempty"`
	// Whether to include all channels in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Broadcasts bool `json:"broadcasts,omitempty"`
	// Whether to include all bots in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Bots bool `json:"bots,omitempty"`
	// Whether to exclude muted chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludeMuted bool `json:"excludeMuted,omitempty"`
	// Whether to exclude read chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludeRead bool `json:"excludeRead,omitempty"`
	// Whether to exclude archived chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludeArchived bool `json:"excludeArchived,omitempty"`
	// Folder¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ID int `json:"ID,omitempty"`
	// Folder¹ name
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Title string `json:"title,omitempty"`
	// Emoji to use as icon for the folder.
	//
	// Use SetEmoticon and GetEmoticon helpers.
	Emoticon string `json:"emoticon,omitempty"`
	// Pinned chats, folders¹ can have unlimited pinned chats
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	PinnedPeers []tg.InputPeerClass `json:"pinnedPeers,omitempty"`
	// Include the following chats in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	IncludePeers []tg.InputPeerClass `json:"includePeers,omitempty"`
	// Exclude the following chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludePeers []tg.InputPeerClass `json:"excludePeers,omitempty"`
}
