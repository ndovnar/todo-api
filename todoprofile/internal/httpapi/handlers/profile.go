package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"todoprofile/internal/converter"
	"todoprofile/internal/dto"
	"todoprofile/internal/service"

	"todolib/auth"
	"todolib/db"
	"todolib/ginhelper"
)

type Profiles struct {
	profileService service.ProfileService
}

func NewProfiles(profileService service.ProfileService) *Profiles {
	return &Profiles{
		profileService: profileService,
	}
}

func (h *Profiles) HandleGetProfile(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	profile, err := h.profileService.GetProfile(ctx, &service.GetProfileParams{
		UserID: claims.UserID,
	})
	if err != nil {
		if err == db.ErrNotFound {
			ctx.Error(ginhelper.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, converter.ProfileModelToDTOResponse(profile))
}

func (h *Profiles) HandleCreateProfile(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	var req dto.CreateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}
	profile, err := h.profileService.CreateProfile(ctx, &service.CreateProfileParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserID:    claims.UserID,
	})
	if err != nil {
		if err == db.ErrDuplicateKey {
			ctx.Error(ginhelper.NewHttpErrorWithDescription(http.StatusUnprocessableEntity, "user already exists"))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, converter.ProfileModelToDTOResponse(profile))
}

func (h *Profiles) HandleUpdateProfile(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	var req dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ginhelper.NewHttpError(http.StatusBadRequest))
		return
	}

	profile, err := h.profileService.UpdateProfile(ctx, &service.UpdateProfileParams{
		UserID:    claims.UserID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		if err == db.ErrNotFound {
			ctx.Error(ginhelper.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.JSON(http.StatusOK, converter.ProfileModelToDTOResponse(profile))
}

func (h *Profiles) HandleDeleteProfile(ctx *gin.Context) {
	claims := auth.GetClaimsFromContext(ctx)

	err := h.profileService.DeleteProfile(ctx, &service.DeleteProfileParams{
		UserID: claims.UserID,
	})
	if err != nil {
		if err == db.ErrNotFound {
			ctx.Error(ginhelper.NewHttpError(http.StatusNotFound))
		} else {
			ctx.Error(ginhelper.NewHttpError(http.StatusInternalServerError))
		}

		return
	}

	ctx.Status(http.StatusOK)
}
